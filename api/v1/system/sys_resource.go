package system

import (
	"fmt"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"sophliteos/global"
	"sophliteos/logger"

	"sophliteos/pkg/constant"
	"sophliteos/pkg/dto"
	"sophliteos/pkg/dto/response"
	"sophliteos/pkg/handle"
	"sophliteos/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type ResourceApi struct{}

func (b *ResourceApi) NewResource(c *gin.Context) {
	switch global.Arch {
	case "arm64":
		resource := GetArmResource(c)
		c.JSON(http.StatusOK, handle.Success(resource))
	default:
		resource := getPcieResource(c)
		c.JSON(http.StatusOK, handle.Success(resource))
	}
}

func getPcieResource(c *gin.Context) response.PcieResource {
	ctrlResource, err := service.GetCtrlResource()
	if err != nil {
		logger.Error("error is :%v", err)
		logger.Error("ctrlResource is :%v", ctrlResource)
	}

	ctrlBasic, deviceSn, err := service.GetCtrlBasic()
	if err != nil {
		logger.Error("error is :%v", err)
		logger.Error("ctrlBasic is :%v", ctrlBasic)
	}

	deviceType := strings.ToUpper(ctrlBasic.Configure.Basic.DeviceType)
	global.DeviceType = deviceType

	global.SSmLists.CoreSsm = []response.SsmVersion{}
	resource := response.PcieResource{
		DeviceSn:        deviceSn,
		DeviceName:      ctrlBasic.Configure.Basic.DeviceName,
		DeviceType:      deviceType,
		SdkVersion:      ctrlBasic.System.SdkVersion,
		OperatingSystem: ctrlBasic.System.OperatingSystem,
		RunTime:         ctrlBasic.System.Runtime,
		BuildTime:       ctrlBasic.System.BuildTime,
		BmssmVersion:    ctrlBasic.System.BmssmVersion,
		Cpu:             getCpu(ctrlResource.CentralProcessingUnit.Cpu),
		Memory:          getMemory(ctrlResource.CentralProcessingUnit.Memory),
		NetCard:         getNetCarts(ctrlResource.CentralProcessingUnit.NetCard),
	}
	for _, disk := range ctrlResource.CentralProcessingUnit.Disk {
		resource.Disk = append(resource.Disk, response.DiskInfo(disk))
	}

	var boards []response.BoardPcie

	for index, board := range ctrlResource.CoreComputingUnit.Board {
		var chips []response.Chip
		var tpuTotal, tpuUsed, memTotal, memUsed float64

		for _, chip := range board.Chip {
			chipIndex, _ := strconv.Atoi(chip.ChipIndex)
			chips = append(chips, response.Chip{
				ChipIndex:                     chipIndex,
				Health:                        chip.Health,
				Temperature:                   chip.Temperature,
				MemoryUsedBytes:               int64(chip.Memory.Total - chip.Memory.Available),
				MemoryTotalBytes:              int64(chip.Memory.Total),
				TpuUtililizationRate:          chip.UtilizationRate,
				TheoretialCalculationCapacity: chip.CalculationCapacity,
				Slot:                          chip.Slot,
				ChipType:                      chip.ChipType,
			})
			chip.CalculationCapacityInt8 = chip.CalculationCapacity
			resource.TPU.Total += chip.CalculationCapacity
			resource.TPU.Used += float64(chip.UtilizationRate) * chip.CalculationCapacity / 100

			tpuTotal += chip.CalculationCapacity
			tpuUsed += float64(chip.UtilizationRate) * chip.CalculationCapacity / 100
			memTotal += chip.Memory.Total
			memUsed += (chip.Memory.Total - chip.Memory.Available)
		}
		if len(board.Chip) == 0 {
			logger.Error("无法读取核心版芯片信息 SN=%s", board.BoardSn)
			continue
		}
		boards = append(boards, response.BoardPcie{
			Number:      index,
			BoardSn:     board.BoardSn,
			BoardType:   board.BoardType,
			Temperature: board.Temperature,
			FanSpeed:    board.FanspeedPercent,
			SdkVersion:  board.SdkVersion,
			Chip:        chips,
			TpuTotal:    tpuTotal,
			TpuUsed:     tpuUsed,
			MemTotal:    memTotal,
			MemUsed:     memUsed,
		})
		resource.TPU.MemTotal += memTotal
		resource.TPU.MemUsed += memUsed
	}

	resource.CoreComputingUnit = response.PcieCoreComputingUnit{
		Board: boards,
	}

	for _, netCard := range ctrlResource.CentralProcessingUnit.NetCard {
		if netCard.IP != "" {
			resource.DeviceIP = netCard.IP
			break
		}
	}

	global.SSmLists.CtrlSsm.DeviceSn = ctrlResource.DeviceSn
	global.SSmLists.CtrlSsm.Ip = resource.DeviceIP
	global.SSmLists.CtrlSsm.Version = "V" + ctrlResource.CentralProcessingUnit.BmssmVersion + "-" + ctrlResource.CentralProcessingUnit.BuildTime
	global.SdkVersion = ctrlBasic.System.SdkVersion

	logger.Debug("resource is :%v", resource)

	return resource
}

func GetArmResource(c *gin.Context) response.Resource {
	ctrlResource, err := service.GetCtrlResource()
	if err != nil {
		logger.Error("error is :%v", err)
		logger.Error("ctrlResource is :%v", ctrlResource)
	}

	ctrlBasic, deviceSn, err := service.GetCtrlBasic()
	if err != nil {
		logger.Error("error is :%v", err)
		logger.Error("ctrlBasic is :%v", ctrlBasic)
	}

	deviceType := strings.ToUpper(ctrlBasic.Configure.Basic.DeviceType)
	switch {
	case strings.Contains(deviceType, "SE5"):
		deviceType = "SE5"
	case strings.Contains(deviceType, "SE6"):
		deviceType = "SE6"
	case strings.Contains(deviceType, "SE7") || deviceType == "BM1684X EVB":
		deviceType = "SE7"
	case strings.Contains(deviceType, "SE8"):
		deviceType = "SE8"
	case strings.Contains(deviceType, "SE9"):
		deviceType = "SE9"
	default:
		deviceType = GetDeviceType()
	}
	global.DeviceType = deviceType

	var emmcDiskCapacity = getDiskCapacity(ctrlResource.CentralProcessingUnit.Disk, true)
	var diskCapacity = getDiskCapacity(ctrlResource.CentralProcessingUnit.Disk, false)

	global.SSmLists.CoreSsm = []response.SsmVersion{}
	resource := response.Resource{
		DeviceSn:        deviceSn,
		CtrlBoardSn:     ctrlBasic.ChipSn,
		DeviceName:      ctrlBasic.Configure.Basic.DeviceName,
		DeviceType:      deviceType,
		SdkVersion:      ctrlBasic.System.SdkVersion,
		OperatingSystem: ctrlBasic.System.OperatingSystem,
		RunTime:         ctrlBasic.System.Runtime,
		BuildTime:       ctrlBasic.System.BuildTime,
		BmssmVersion:    "V" + ctrlBasic.System.BmssmVersion + "-" + ctrlBasic.System.BuildTime,
		Cpu:             getCpu(ctrlResource.CentralProcessingUnit.Cpu),
		Memory:          getMemory(ctrlResource.CentralProcessingUnit.Memory),
		Disk:            getDisks(ctrlResource.CentralProcessingUnit.Disk),
		NetCard:         getNetCarts(ctrlResource.CentralProcessingUnit.NetCard),
		Int8Count: response.ResourceCount{
			Unit: "TOPS",
		},
		Fp16Count: response.ResourceCount{
			Unit: "TFLOPS",
		},
		Fp32Count: response.ResourceCount{
			Unit: "TFLOPS",
		},
		CpuCount: response.ResourceCount{
			Available: ctrlResource.CentralProcessingUnit.Cpu.Cores,
			Total:     ctrlResource.CentralProcessingUnit.Cpu.Cores,
			Unit:      "核",
			Desc:      fmt.Sprintf("@%dMHz", ctrlResource.CentralProcessingUnit.Cpu.Frequency),
		},
		MemoryCount: response.ResourceCount{
			Available: ctrlResource.CentralProcessingUnit.Memory.Available,
			Total:     ctrlResource.CentralProcessingUnit.Memory.Total,
			Unit:      "GB",
			Desc:      "LPDDR4x",
		},
		EMMCCount: response.ResourceCount{
			Available: emmcDiskCapacity,
			Total:     emmcDiskCapacity,
			Unit:      "GB",
		},
		DiskCount: response.ResourceCount{
			Available: diskCapacity,
			Total:     diskCapacity,
			Unit:      "GB",
		},
	}

	global.SE6Cores = nil

	var boards []response.Board

	for _, board := range ctrlResource.CoreComputingUnit.Board {
		cpu := getCpu(board.CoreSys.Cpu)
		memory := getMemory(board.CoreSys.Mem)
		disk := getDisks(board.CoreSys.Disks)
		netCart := getNetCarts(board.CoreSys.NetCards)
		var chips []response.Chip
		coreSys := board.CoreSys

		resource.MemoryCount.Available += coreSys.Mem.Available
		resource.MemoryCount.Total += coreSys.Mem.Total

		chipIndex, _ := strconv.Atoi(board.Chip[0].ChipIndex)

		for _, chip := range board.Chip {
			chips = append(chips, response.Chip{
				ChipIndex:                     chipIndex,
				Health:                        chip.Health,
				Temperature:                   chip.Temperature,
				MemoryUsedBytes:               int64(chip.Memory.Total - chip.Memory.Available),
				MemoryTotalBytes:              int64(chip.Memory.Total),
				TpuUtililizationRate:          chip.UtilizationRate,
				TheoretialCalculationCapacity: chip.CalculationCapacity,
				Deploys:                       []response.Deploy{},
				ChipType:                      chip.ChipType,
			})
			chip.CalculationCapacityInt8 = chip.CalculationCapacity
			if chip.Health == 0 { // 正常
				resource.Int8Count.Health++
				resource.Int8Count.Available += chip.CalculationCapacityInt8

				resource.Fp16Count.Health++
				resource.Fp16Count.Available += chip.CalculationCapacityFp16

				resource.Fp32Count.Health++
				resource.Fp32Count.Available += chip.CalculationCapacityFp32

				resource.CpuCount.Health++
				resource.CpuCount.Available += cpu.Cores

				resource.MemoryCount.Health++
				resource.MemoryCount.Available += chip.Memory.Available
				resource.MemoryCount.Total += chip.Memory.Total

				resource.DiskCount.Health++
				resource.DiskCount.Available += getDiskCapacity(coreSys.Disks, false)

				resource.EMMCCount.Health++
				resource.EMMCCount.Available += getDiskCapacity(coreSys.Disks, true)
			} else {
				resource.Int8Count.UnHealth++
				resource.Fp16Count.UnHealth++
				resource.Fp32Count.UnHealth++
				resource.CpuCount.UnHealth++
				resource.MemoryCount.UnHealth++
				resource.DiskCount.UnHealth++
			}
			resource.Int8Count.Total += chip.CalculationCapacityInt8
			resource.Fp16Count.Total += chip.CalculationCapacityFp16
			resource.Fp32Count.Total += chip.CalculationCapacityFp32
			resource.CpuCount.Total += cpu.Cores

			resource.Fp32Count.Total = resource.Int8Count.Total / 8
			resource.Fp16Count.Total = resource.Int8Count.Total / 4

			resource.DiskCount.Total += getDiskCapacity(coreSys.Disks, false)
			resource.EMMCCount.Total += getDiskCapacity(coreSys.Disks, true)

		}
		if len(board.Chip) == 0 {
			logger.Error("无法读取核心版芯片信息 SN=%s", board.BoardSn)
			continue
		}

		var coreIp string
		if deviceType == "SE6" || deviceType == "SE8" {
			if len(netCart) > 0 && len(netCart[0].Ip) > 0 {
				coreIp = netCart[0].Ip
				chipIndex = constant.NumberMap[extractIP(coreIp)]
			}
			// 核心板SSM版本信息
			sv := response.SsmVersion{
				ChipIndex: chipIndex,
				DeviceSn:  board.BoardSn,
				Ip:        coreIp,
				Version:   "V" + coreSys.BmssmVersion + "-" + coreSys.BuildTime,
			}
			global.SSmLists.CoreSsm = append(global.SSmLists.CoreSsm, sv)

			// 核心板SSH IP
			global.SE6Cores = append(global.SE6Cores, global.SE6Core{
				Ip:       coreIp,
				DeviceSn: board.BoardSn,
				Number:   chipIndex,
			})
		}

		boards = append(boards, response.Board{
			Cpu:         cpu,
			Memory:      memory,
			Disk:        disk,
			NetCard:     netCart,
			Number:      chipIndex,
			BoardSn:     board.BoardSn,
			BoardType:   board.BoardType,
			Temperature: board.Temperature,
			FanSpeed:    board.FanspeedPercent,
			SdkVersion:  board.SdkVersion,
			Chip:        chips,
			Ip:          coreIp,
			SSMVersion:  "V" + coreSys.BmssmVersion + "-" + coreSys.BuildTime,
		})
	}

	resource.MemoryCount.Available = resource.MemoryCount.Available / 1024
	resource.MemoryCount.Total = resource.MemoryCount.Total / 1024
	resource.CoreComputingUnit = response.CoreComputingUnit{
		Board: boards,
	}

	switch deviceType {
	case "SE5", "SE7", "SE9":
		for _, netCard := range ctrlResource.CentralProcessingUnit.NetCard {
			if netCard.NetCardName == "eth0" {
				resource.WanIP = netCard.IP
			} else if netCard.NetCardName == "eth1" {
				resource.LanIP = netCard.IP
			}
		}
		global.SSmLists.CoreSsm = []response.SsmVersion{}
	case "SE6", "SE8":
		for _, netCard := range ctrlResource.CentralProcessingUnit.NetCard {
			switch netCard.NetCardName {
			case "eth0", "eth1":
				if netCard.IP != "" {
					resource.LanIP = resource.LanIP + "," + netCard.IP
				}
			case "enp3s0":
				if netCard.IP != "" {
					resource.WanIP = netCard.IP
				}
			case "enp4s0", "enp7s0":
				if netCard.IP != "" {
					resource.WanIP = resource.WanIP + "," + netCard.IP
				}
			case "enp8s0":
				if netCard.IP != "" {
					resource.LanIP = resource.LanIP + "," + netCard.IP
				}
			}
		}
		// SE6设备核心板排序
		sort.Slice(global.SE6Cores, func(i, j int) bool {
			return global.SE6Cores[i].Number < global.SE6Cores[j].Number
		})
		sort.Slice(global.SSmLists.CoreSsm, func(i, j int) bool {
			return global.SSmLists.CoreSsm[i].ChipIndex < global.SSmLists.CoreSsm[j].ChipIndex
		})
	}

	if len(resource.LanIP) > 1 && resource.LanIP[0] == ',' {
		resource.LanIP = resource.LanIP[1:]
	}
	if len(resource.WanIP) > 1 && resource.WanIP[0] == ',' {
		resource.WanIP = resource.WanIP[1:]
	}
	resource.DeviceIP = resource.WanIP

	global.SSmLists.CtrlSsm.DeviceSn = ctrlResource.DeviceSn
	global.SSmLists.CtrlSsm.Ip = resource.DeviceIP
	global.SSmLists.CtrlSsm.Version = "V" + ctrlResource.CentralProcessingUnit.BmssmVersion + "-" + ctrlResource.CentralProcessingUnit.BuildTime
	global.SdkVersion = ctrlBasic.System.SdkVersion
	// 兼容3wan口se6
	if deviceType == "SE6" {
		resource.Fp16Count.Total = 0

		result, _ := service.GetIP()
		var se6Ips []dto.Ip
		err = mapstructure.Decode(result.Result, &se6Ips)
		handle.HandleError(err)
		// ssm未返回，兼容添加网卡
		appendNetCard(se6Ips, resource.NetCard)
	}

	logger.Debug("resource is :%v", resource)

	return resource
}

func appendNetCard(ips []dto.Ip, netCards []response.NetCard) {
	for _, ip := range ips {
		if containsNetCard(netCards, ip.NetCardName) {
			continue
		}
		netCards = append(netCards, response.NetCard{
			Ip:        ip.IP,
			Name:      ip.Name,
			Bandwidth: ip.Bandwidth,
			Mac:       ip.Mac,
			NetIn:     ip.NetRx,
			NetOut:    ip.NetTx,
		})
	}
}

func containsNetCard(netCards []response.NetCard, name string) bool {
	for _, item := range netCards {
		if item.Name == name {
			return true
		}
	}
	return false
}

func getCpu(cpu dto.CPU) response.CPU {
	return response.CPU{
		Cores:     cpu.Cores,
		Frequency: cpu.Frequency,
		Usage:     cpu.UtilizationRate,
		Type:      cpu.Type,
		Arch:      cpu.Arch,
	}
}

func getMemory(memory dto.Memory) response.Memory {
	if memory.Total == 0 {
		return response.Memory{
			Total: 0,
			Usage: 0,
		}
	}
	return response.Memory{
		Total: memory.Total,
		Usage: 100 - (memory.Available/memory.Total)*100,
	}
}

func getDisks(disk []dto.Disk) []response.Disk {
	var disks []response.Disk

	free, total := addDisk(disk, true)
	if math.Abs(total) < 0.1 {
		total = 1
		free = 0
	}
	disks = append(disks, response.Disk{
		ID:    "mmcblk",
		Total: total,
		Usage: (1 - float64(free/total)) * 100,
	})

	free, total = addDisk(disk, false)
	if math.Abs(total) < 0.1 {
		total = 1
		free = 0
	}
	dsk := response.Disk{
		ID:    "others",
		Total: total,
		Usage: (1 - float64(free/total)) * 100,
	}
	if dsk.Total == 1 && dsk.Usage == 100 {
		dsk.Total = 0.0
		dsk.Usage = 0.0
	}
	disks = append(disks, dsk)
	return disks
}

func addDisk(disk []dto.Disk, eMMC bool) (float64, float64) {
	free := 0.0
	total := 0.0

	for _, _disk := range disk {
		if _disk.Total < 1 {
			continue
		}
		if eMMC && strings.HasPrefix(_disk.DiskName, "/dev/mmc") {
			total = total + _disk.Total
			free += _disk.Free
		} else if !eMMC && !strings.HasPrefix(_disk.DiskName, "/dev/mmc") {
			total += _disk.Total
			free += _disk.Free
		}
	}
	return free, total
}

func getDiskCapacity(disk []dto.Disk, eMMC bool) float64 {
	var capacity float64
	for _, _disk := range disk {
		if _disk.Total == 0 {
			continue
		}
		if eMMC && strings.HasPrefix(_disk.DiskName, "/dev/mmc") {
			capacity += _disk.Total
		} else if !eMMC && !strings.HasPrefix(_disk.DiskName, "/dev/mmc") {
			capacity += _disk.Total
		}
	}
	return capacity / 1024
}

func getNetCarts(netCards []dto.NetCard) []response.NetCard {
	var res []response.NetCard
	for _, netCard := range netCards {
		if !strings.HasPrefix(netCard.NetCardName, "en") && !strings.HasPrefix(netCard.NetCardName, "eth") {
			continue
		}
		res = append(res, response.NetCard{
			Ip:        netCard.IP,
			Name:      netCard.NetCardName,
			Bandwidth: netCard.Bandwidth,
			Mac:       netCard.Mac,
			NetIn:     netCard.NetRx,
			NetOut:    netCard.NetTx,
		})
	}
	return res
}

func extractIP(ipAddress string) string {
	parts := strings.Split(ipAddress, ".")
	if len(parts) == 4 {
		res := parts[2] + "." + parts[3]
		return res
	}
	return ""
}
