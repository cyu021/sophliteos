package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sophliteos/logger"
	"sophliteos/pkg/utils/cmd"
	"sophliteos/pkg/utils/httpclient"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	containerRes "sophliteos/pkg/dto/response"

	"github.com/spf13/viper"
)

var imageId int = 0

type ContainerService struct {
	isInit           bool
	container_type   string
	containerCfgLock sync.RWMutex
	containerCfgPath string
	registryRe       *regexp.Regexp
	pullingInfo      map[string]containerRes.PullingInfo
}

func (self *ContainerService) Init() {
	self.isInit = false
	self.container_type = DOCKER
	self.containerCfgPath = "/etc/docker/daemon.json"
	self.registryRe = regexp.MustCompile("^((\\d|[1-9]\\d|1\\d{2}|2[0-4]\\d|25[0-5])\\.){3}(\\d|[1-9]\\d|1\\d{2}|2[0-4]\\d|25[0-5]):(6553[0-5]|655[0-2]\\d|65[0-4]\\d{2}|6[0-4]\\d{3}|[1-5]\\d{4}|[1-9]\\d{0,3})$")
	self.pullingInfo = make(map[string]containerRes.PullingInfo)

	cpuCmd := "lscpu"
	outStr, errStr, err := cmd.Output(cpuCmd)
	if errStr != "" {
		logger.Error("Failed to get cpu info, error: %s", errStr)
		return
	}
	if err != nil {
		logger.Error("Failed to run command %s, %s: %s", cpuCmd, err, errStr)
		return
	}

	lines := strings.Split(outStr, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Architecture:") {
			arch := strings.TrimSpace(line[len("Architecture:"):])
			if arch == "riscv64" {
				outStr, errStr, err = cmd.Output("podman -v")
				if err == nil && len(outStr) > 0 {
					self.container_type = PODMAN
					self.containerCfgPath = "/etc/containers/registries.conf"
				}
			}
			break
		}
	}
	self.isInit = true
	go self.ListImages()
}

func (self *ContainerService) GetContainerType() string {
	self.containerCfgLock.Lock()
	if !self.isInit {
		self.Init()
	}
	self.containerCfgLock.Unlock()
	return self.container_type
}

func (self *ContainerService) ListInsecureRegistries() ([]string, error) {
	if self.GetContainerType() == DOCKER {
		return self.getDockerInsecureRegistries()
	} else if self.GetContainerType() == PODMAN {
		return self.getPodmanInsecureRegistries()
	} else {
		err := fmt.Errorf("container_type unset")
		logger.Error(err.Error())
		return nil, err
	}
}

func (self *ContainerService) AddInsecureRegistries(reg string) error {
	if self.GetContainerType() == DOCKER {
		return self.AddDockerInsecureRegistries(reg)
	} else if self.GetContainerType() == PODMAN {
		return self.AddPodmanInsecureRegistries(reg)
	} else {
		err := fmt.Errorf("container_type unset")
		logger.Error(err.Error())
		return err
	}
}

func (self *ContainerService) RemoveInsecureRegistries(reg string) error {
	if self.GetContainerType() == DOCKER {
		return self.RemoveDockerInsecureRegistries(reg)
	} else if self.GetContainerType() == PODMAN {
		return self.RemovePodmanInsecureRegistries(reg)
	} else {
		err := fmt.Errorf("container_type unset")
		logger.Error(err.Error())
		return err
	}
}

func (self *ContainerService) ListInsecureImageTags(reg string) (containerRes.RegistryImages, error) {
	var images containerRes.RegistryImages
	var res containerRes.RegistryImages
	imagesUrl := reg + "/v2/_catalog"

	imagesByte, err := httpclient.NewRequest(imagesUrl, "GET", nil, nil)

	if err != nil {
		logger.Error("Error send request to %s", imagesUrl)
		return res, err
	}
	err = json.Unmarshal(imagesByte, &images)
	if err != nil {
		return res, err
	}

	for i, image := range images.Repositories {
		var tags containerRes.RegistryImageTags
		tagsUrl := reg + "/v2/" + image + "/tags/list"
		tagsByte, err := httpclient.NewRequest(tagsUrl, "GET", nil, nil)
		if err != nil {
			logger.Error("Error send request to %s", tagsUrl)
			res.Repositories = append(res.Repositories[:i], res.Repositories[i+1:]...)
			continue
		}

		err = json.Unmarshal(tagsByte, &tags)
		if err != nil {
			continue
		}

		for _, tag := range tags.Tags {
			res.Repositories = append(res.Repositories, image+":"+tag)
		}
	}

	return res, nil

}

func (self *ContainerService) ListInsecureImages(reg string) (containerRes.RegistryImages, error) {
	var res containerRes.RegistryImages
	imagesUrl := reg + "/v2/_catalog"

	imagesByte, err := httpclient.NewRequest(imagesUrl, "GET", nil, nil)

	if err != nil {
		logger.Error("Error send request to %s", imagesUrl)
		return res, err
	}
	err = json.Unmarshal(imagesByte, &res)
	return res, err

}

func (self *ContainerService) ListInsecureTags(reg string, image string) (containerRes.RegistryImageTags, error) {
	var res containerRes.RegistryImageTags
	imagesUrl := reg + "/v2/" + image + "/tags/list"
	imagesByte, err := httpclient.NewRequest(imagesUrl, "GET", nil, nil)

	if err != nil {
		logger.Error("Error send request to %s", imagesUrl)
		return res, err
	}
	err = json.Unmarshal(imagesByte, &res)
	fmt.Println(res)

	return res, err

}

func (self *ContainerService) getDockerInsecureRegistries() ([]string, error) {
	var registryList []string = []string{}

	self.containerCfgLock.RLock()
	defer self.containerCfgLock.RUnlock()

	_, err := os.Stat(self.containerCfgPath)
	if err != nil {
		return nil, nil
	}
	byteValue, err := os.ReadFile(self.containerCfgPath)
	if err != nil {
		logger.Error("Failed to read %s config file, error: %s", self.GetContainerType(), err)
		return registryList, nil
	}

	var result map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		logger.Error("Failed to unmarshal docker config, error: %s", err)
		return registryList, err
	}

	if _, exists := result["insecure-registries"]; exists {
		registries, _ := result["insecure-registries"].([]interface{})
		for _, registry := range registries {
			r, _ := registry.(string)
			registryList = append(registryList, r)
		}
	}

	return registryList, nil
}

func (self *ContainerService) AddDockerInsecureRegistries(reg string) error {
	var registryList []string = []string{}
	var f *os.File
	var err, err1 error
	self.containerCfgLock.Lock()
	defer self.containerCfgLock.Unlock()
	_, err = os.Stat(self.containerCfgPath)
	if err != nil {
		f, err1 = os.Create(self.containerCfgPath)
		if err1 != nil {
			logger.Error("Failed to create config file, error: %s", err)
			return err
		}
	}
	byteValue, err1 := os.ReadFile(self.containerCfgPath)
	if err1 != nil {
		logger.Error("Failed to read docker config file, error: %s", err)
		return err1
	}
	if err != nil || len(byteValue) == 0 {
		var address = make(map[string]interface{})
		var varp []string
		varp = append(varp, reg)
		address["insecure-registries"] = varp
		addBytes, err := json.Marshal(address)
		if err != nil {
			logger.Error("Failed to marshal config address to bytes, error: %s", err)
			return err
		} else {
			_, err := f.Write(addBytes)
			if err != nil {
				logger.Error("Failed to write bytes to file %s, error: %s", self.containerCfgPath, err)
				return err
			}
		}
		f.Close()
	} else {
		var result map[string]interface{}
		err = json.Unmarshal(byteValue, &result)
		if err != nil {
			logger.Error("Failed to unmarshal docker config, error: %s", err)
			return err
		}

		if _, exists := result["insecure-registries"]; exists {
			registries, _ := result["insecure-registries"].([]interface{})
			for _, registry := range registries {
				r, _ := registry.(string)
				if r == reg {
					logger.Info("仓库 %s 已存在", reg)
					return fmt.Errorf("仓库 %s 已存在", reg)
				}
				registryList = append(registryList, r)
			}
		}

		registryList = append(registryList, reg)
		result["insecure-registries"] = registryList

		byteValue, err = json.Marshal(result)
		if err != nil {
			logger.Error("Failed to marshal docker config, error: %s", err)
			return err
		}

		err = os.WriteFile(self.containerCfgPath, byteValue, 0755)
		if err != nil {
			logger.Error("Failed to write docker config file, error: %s", err)
			return err
		}
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		logger.Error("Failed to reload daemon, error: %s", err)
		return err
	}

	cmd = exec.Command("systemctl", "reload", "docker")
	err = cmd.Run()
	if err != nil {
		logger.Error("Failed to restart docker service, error: %s", err)
		return err
	}

	return nil
}

func (self *ContainerService) RemoveDockerInsecureRegistries(reg string) error {
	var registryList []string = []string{}
	var err, err1 error
	self.containerCfgLock.Lock()
	defer self.containerCfgLock.Unlock()
	_, err = os.Stat(self.containerCfgPath)
	if err != nil {
		logger.Error("no such insecure: %s", reg)
		return err
	}

	byteValue, err1 := os.ReadFile(self.containerCfgPath)
	if err1 != nil {
		logger.Error("Failed to read docker config file, error: %s", err)
		return err1
	}

	if len(byteValue) == 0 {
		logger.Error("Failed to read docker config file, error: %s", err)
		return err1
	} else {
		var result map[string]interface{}
		err = json.Unmarshal(byteValue, &result)
		if err != nil {
			logger.Error("Failed to unmarshal docker config, error: %s", err)
			return err
		}

		isRemoved := false
		if _, exists := result["insecure-registries"]; exists {
			registries, _ := result["insecure-registries"].([]interface{})
			for _, registry := range registries {
				r, _ := registry.(string)
				if r == reg {
					isRemoved = true
					continue
				}
				registryList = append(registryList, r)
			}
		}

		if !isRemoved {
			err := fmt.Errorf("no such insecure: %s", reg)
			logger.Error(err.Error())
			return err
		}
		result["insecure-registries"] = registryList

		byteValue, err = json.Marshal(result)
		if err != nil {
			logger.Error("Failed to marshal docker config, error: %s", err)
			return err
		}

		err = os.WriteFile(self.containerCfgPath, byteValue, 0755)
		if err != nil {
			logger.Error("Failed to write docker config file, error: %s", err)
			return err
		}
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		logger.Error("Failed to reload daemon, error: %s", err)
		return err
	}

	cmd = exec.Command("systemctl", "reload", "docker")
	err = cmd.Run()
	if err != nil {
		logger.Error("Failed to restart docker service, error: %s", err)
		return err
	}

	return nil
}
func (self *ContainerService) getPodmanInsecureRegistries() ([]string, error) {
	self.containerCfgLock.RLock()
	defer self.containerCfgLock.RUnlock()

	v := viper.New()
	v.SetConfigType("toml")
	v.AddConfigPath(filepath.Dir(self.containerCfgPath))
	v.SetConfigName(filepath.Base(self.containerCfgPath))
	_, err := os.Stat(self.containerCfgPath)
	if os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(self.containerCfgPath), os.ModePerm)
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var config map[string][]interface{}

	err = v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	var res []string
	for _, registry := range config["registry"] {
		if regisMap, ok := registry.(map[string]interface{}); ok {

			if insecure, ok := regisMap["insecure"]; ok {
				if insecure == false {
					continue
				}
			}

			if location, ok := regisMap["location"].(string); ok {
				res = append(res, location)
			}
		}
	}
	return res, nil
}

func (self *ContainerService) AddPodmanInsecureRegistries(reg string) error {
	self.containerCfgLock.RLock()
	defer self.containerCfgLock.RUnlock()

	v := viper.New()
	v.SetConfigType("toml")
	v.AddConfigPath(filepath.Dir(self.containerCfgPath))
	v.SetConfigName(filepath.Base(self.containerCfgPath))

	_, err := os.Stat(self.containerCfgPath)
	config := make(map[string][]interface{})
	regis := make(map[string]interface{})
	regis["location"] = reg
	regis["insecure"] = true

	if os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(self.containerCfgPath), os.ModePerm)
		v.Set("registry", append(config["registry"], regis))

	} else if err == nil {
		if err := v.ReadInConfig(); err != nil {
			return err
		}
		var config map[string][]interface{}

		if err := v.Unmarshal(&config); err != nil {
			return err
		}

		for _, registry := range config["registry"] {
			if regisMap, ok := registry.(map[string]interface{}); ok {
				if insecure, ok := regisMap["insecure"]; ok {
					if insecure == false {
						continue
					}
				}
				if location, ok := regisMap["location"]; ok {
					if location == reg {
						return fmt.Errorf("registry %s is existed", reg)
					}
				}
			}
		}
		v.Set("registry", append(config["registry"], regis))

	} else if err != nil {
		return err
	}
	tempConf := filepath.Dir(self.containerCfgPath) + "/newConf.toml"
	v.WriteConfigAs(tempConf)

	cmd := exec.Command("mv", tempConf, self.containerCfgPath)
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		logger.Error("Failed to reload daemon, error: %s", err)
		return err
	}

	return nil
}

func (self *ContainerService) RemovePodmanInsecureRegistries(reg string) error {
	self.containerCfgLock.RLock()
	defer self.containerCfgLock.RUnlock()

	v := viper.New()
	v.SetConfigType("toml")
	v.AddConfigPath(filepath.Dir(self.containerCfgPath))
	v.SetConfigName(filepath.Base(self.containerCfgPath))

	_, err := os.Stat(self.containerCfgPath)
	regis := make(map[string]interface{})
	regis["location"] = reg
	regis["insecure"] = true

	if os.IsNotExist(err) {
		logger.Error("no such insecure: %s", reg)
		return err

	} else if err == nil {
		if err := v.ReadInConfig(); err != nil {
			return err
		}
		var config map[string][]interface{}

		if err := v.Unmarshal(&config); err != nil {
			return err
		}

		isRemoved := false
		for idx, registry := range config["registry"] {
			if regisMap, ok := registry.(map[string]interface{}); ok {
				if insecure, ok := regisMap["insecure"]; ok {
					if insecure == false {
						continue
					}
				}
				if location, ok := regisMap["location"]; ok {
					if location == reg {
						config["registry"] = append(config["registry"][:idx], config["registry"][idx+1:]...)
						isRemoved = true
					}
				}
			}
		}
		v.Set("registry", append(config["registry"], regis))
		if isRemoved {
			tempConf := filepath.Dir(self.containerCfgPath) + "/newConf.toml"
			v.WriteConfigAs(tempConf)

			cmd := exec.Command("mv", tempConf, self.containerCfgPath)
			err = cmd.Run()
			if err != nil {
				return err
			}

			cmd = exec.Command("systemctl", "daemon-reload")
			err = cmd.Run()
			if err != nil {
				logger.Error("Failed to reload daemon, error: %s", err)
				return err
			}

			return nil
		}
		err := fmt.Errorf("no such insecure: %s", reg)
		logger.Error(err.Error())
		return err

	}
	return err
}

func (self *ContainerService) checkRegistry(image string) bool {
	idx := strings.Index(image, "/")
	if idx == -1 {
		return false
	}

	registry := image[:idx]
	if self.registryRe.MatchString(registry) {
		return true
	}
	return false
}

func (self *ContainerService) PullImage(img string) error {
	buf := fmt.Sprintf("sudo %s images -q %s", self.GetContainerType(), img)
	out, err := exec.Command("/bin/sh", "-c", buf).Output()
	if err == nil && len(out) > 0 {
		return nil
	}

	go self.doPullImage(img)
	return nil
}

func (self *ContainerService) CheckImagePulled(img string) int {
	if _, ok := self.pullingInfo[img]; ok {
		return 1
	}
	if _, err := self.InspectImage(img); err == nil {
		return 0
	} else {
		return 2
	}
}

func (self *ContainerService) doPullImage(img string) error {
	pullCmd := fmt.Sprintf("sudo %s pull %s", self.GetContainerType(), img)
	imageId++
	self.pullingInfo[img] = containerRes.PullingInfo{
		Name:   img,
		Cmd:    pullCmd,
		Status: 1,
		Id:     imageId,
	}

	_, errStr, err := cmd.Output(pullCmd)

	if tmp, ok := self.pullingInfo[img]; ok {
		if err == nil && errStr == "" {
			tmp.Status = 0
			self.pullingInfo[img] = tmp
		} else {
			logger.Error("pull image failed, cmd: " + pullCmd + " errstr: " + errStr)
			tmp.Status = 2
			self.pullingInfo[img] = tmp
		}
	} else {
		logger.Error("bad error")
	}
	return err
}

func (self *ContainerService) ListPulling() ([]containerRes.PullingInfo, error) {
	var res []containerRes.PullingInfo
	for _, info := range self.pullingInfo {
		res = append(res, info)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Id < res[j].Id
	})
	return res, nil
}

func (self *ContainerService) RemoveContainer(app string) error {
	err := self.StopContainer(app, 30*time.Second)
	if err != nil {
		logger.Error("Error stop container")
	}

	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo %s rm %s", self.GetContainerType(), app))
	out, err := cmd.Output()
	if err != nil {
		logger.Error("Error remove container, cmd: ", cmd, ", err: ", err.Error())
		logger.Error("output: ", out)
	}
	return err
}

func (self *ContainerService) RemoveImage(img string) error {

	// cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo %s rmi %s", self.GetContainerType(), img))
	rmCmd := fmt.Sprintf("sudo %s rmi %s", self.GetContainerType(), img)
	outStr, errStr, err := cmd.Output(rmCmd)

	if err != nil {
		logger.Errorln("Error remove container, cmd: ", rmCmd, ", err: ", err.Error())
		logger.Errorln("output: ", outStr, " errStr: ", errStr)
		if strings.Contains(errStr, "image is being used") {
			return fmt.Errorf("镜像正在使用，无法删除")
		}
	}
	return err
}
func (self *ContainerService) TagImage(img string, tag_img string) error {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo %s tag %s %s", self.GetContainerType(), img, tag_img))
	out, err := cmd.Output()
	if err != nil {
		logger.Error("Error remove container, cmd: ", cmd, ", err: ", err.Error())
		logger.Error("output: ", out)
	}
	return err
}

func (self *ContainerService) RestartContainer(app string, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", fmt.Sprintf("sudo %s restart %s", self.GetContainerType(), app))
	err := cmd.Run()
	if ctx.Err() != nil {
		err = fmt.Errorf("重启容器超时")
		logger.Error(err.Error())
		return err
	}
	if err != nil {
		logger.Error("Error restart container, cmd: ", cmd)
	}
	return err
}

func (self *ContainerService) StopContainer(app string, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", fmt.Sprintf("sudo %s stop %s", self.GetContainerType(), app))
	err := cmd.Run()
	if ctx.Err() != nil {
		logger.Error("stop container timeout, cmd: ", cmd)
		return ctx.Err()
	}
	if err != nil {
		logger.Error("Error stop container, cmd: ", cmd)
	}
	return err
}

func (self *ContainerService) StartContainer(app string, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", fmt.Sprintf("sudo %s start %s", self.GetContainerType(), app))
	err := cmd.Run()
	if ctx.Err() != nil {
		logger.Error("start container timeout, cmd: ", cmd)
		return ctx.Err()
	}
	if err != nil {
		logger.Error("Error start container, cmd: ", cmd)
	}
	return err
}

func (self *ContainerService) InspectContainer(app string) (containerRes.InspectJson, error) {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo %s inspect --format='{{json .State}}' %s", self.GetContainerType(), app))
	out, err := cmd.Output()
	var res containerRes.InspectJson
	if err != nil {
		logger.Error("Error inspect container, cmd: ", cmd)
		return res, err
	}
	err = json.Unmarshal(out, &res)
	if err != nil {
		logger.Error("Error unmarshal inspect output, cmd: ", cmd)
	}
	return res, err
}

func (self *ContainerService) InspectImage(app string) (containerRes.InspectImageJson, error) {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo %s inspect --format='{{json .}}' %s", self.GetContainerType(), app))
	out, err := cmd.Output()
	var res containerRes.InspectImageJson
	if err != nil {
		logger.Error("Error inspect container, cmd: %s", cmd)
		return res, err
	}
	err = json.Unmarshal(out, &res)
	if err != nil {
		logger.Error("Error unmarshal inspect output, cmd: %s", cmd)
	}
	return res, err
}

func (self *ContainerService) ListImages() ([]containerRes.ImageInfo, error) {
	var res []containerRes.ImageInfo

	out, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo %s images --format='{{json .}}'", self.GetContainerType())).Output()
	if err != nil {
		return nil, err
	}

	if len(out) == 0 {
		return nil, nil
	}

	out = bytes.TrimRight(out, "\n")
	lines := bytes.Split(out, []byte("\n"))
	for index, i := range lines {
		var img containerRes.ImageInfo
		switch self.GetContainerType() {
		case PODMAN:
			err := json.Unmarshal(i, &img)
			if err != nil || img.Repo == "" {
				logger.Errorln("Error list podman images index: ", index, err)
				continue
			}
			img.Size /= 1000000

		case DOCKER:
			var dockerImg containerRes.DockerImageInfo
			err := json.Unmarshal(i, &dockerImg)
			if err != nil || dockerImg.Repo == "" {
				logger.Errorln("Error list docker images index: ", index, err)
				continue
			}
			timeLayout := "2006-01-02 15:04:05 -0700 MST"
			t, err := time.Parse(timeLayout, dockerImg.CreatedAt)
			if err != nil {
				logger.Error(err.Error())
				t, _ = time.Parse(timeLayout, timeLayout)
			}
			img.Created = t.Unix()
			img.Repo = dockerImg.Repo
			img.Tag = dockerImg.Tag
			img.ImageId = dockerImg.ImageId
			img.Size = 0

			re := regexp.MustCompile(`(\d+(\.\d+)?)`)
			match := re.FindStringSubmatch(dockerImg.Size)
			if len(match) > 1 {
				sizeFloat, err := strconv.ParseFloat(match[0], 64)
				if err == nil {
					unit := dockerImg.Size[len(dockerImg.Size)-2:]
					img.Size = sizeFloat
					if unit == "kB" {
						img.Size /= 1000
					} else if unit == "GB" {
						img.Size *= 1000
					}
					break
				}
				logger.Errorln("docker image size format error: ", dockerImg.Size, err.Error())
				break
			}
			logger.Errorln("docker image size format error: ", dockerImg.Size)

		default:
			continue
		}

		res = append(res, img)
	}

	return res, nil

}

func (self *ContainerService) ListContainers() ([]containerRes.ListContainerJson, error) {
	containerCmd := fmt.Sprintf("sudo %s ps -a --format '{{.ID}}|{{.Image}}|{{.Names}}|{{.CreatedAt}}|{{.Status}}|{{.Ports}}'", self.GetContainerType())
	outStr, errStr, err := cmd.Output(containerCmd)
	if errStr != "" {
		logger.Error("Failed to get container info, error: %s", errStr)
		return nil, fmt.Errorf(errStr)
	}
	if err != nil {
		logger.Error("Failed to run command %s, %s: %s", containerCmd, err, errStr)
		return nil, err
	}
	if outStr == "" {
		return nil, nil
	}
	outStr = strings.TrimRight(outStr, "\n")
	lines := strings.Split(outStr, "\n")

	var res []containerRes.ListContainerJson
	for _, line := range lines {
		attrs := strings.Split(line, "|")
		createdAt := attrs[3]

		if t, err := time.Parse("2006-01-02 15:04:05 +0800 CST", createdAt); err == nil {
			createdAt = t.Format("2006/01/02 15:04:05")
		}

		// if t, err := time.Parse(time.RFC3339Nano, containerInfo.FinishedAt); err == nil {
		// 	containerInfo.FinishedAt = t.Format("2006/01/02 15:04:05")
		// }
		var ports []string
		if len(attrs[5]) != 0 {
			ports = strings.Split(attrs[5], ", ")

		}

		isRunning := false
		if attrs[4][0:2] == "Up" {
			isRunning = true
		}
		res = append(res, containerRes.ListContainerJson{
			Name:      attrs[2],
			Image:     attrs[1],
			ID:        attrs[0],
			Status:    attrs[4],
			CreatedAt: createdAt,
			Ports:     ports,
			Running:   isRunning,
		})
	}
	return res, nil

}
