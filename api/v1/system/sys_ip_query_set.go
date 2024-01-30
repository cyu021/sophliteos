package system

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"sophliteos/global"
	"sophliteos/logger"

	"sophliteos/pkg/buserr"
	"sophliteos/pkg/dto"
	"sophliteos/pkg/dto/response"
	"sophliteos/pkg/handle"
	"sophliteos/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type IpApi struct{}

func (b *IpApi) IpQuery(c *gin.Context) {
	result, _ := service.GetIP()
	var Ips []dto.Ip
	err := mapstructure.Decode(result.Result, &Ips)
	handle.HandleError(err)

	var NewIps []response.NetIp
	if global.DeviceType == "SE6" || global.DeviceType == "SE8" {
		for _, ip := range Ips {
			if !strings.HasPrefix(ip.NetCardName, "en") {
				continue
			}
			var dns string
			if len(ip.DNS) > 0 {
				dns = ip.DNS[0]
			} else {
				dns = ""
			}

			netIp := response.NetIp{
				NetCardName: ip.NetCardName,
				Bandwidth:   ip.Bandwidth,
				DeltaRx:     ip.DeltaRx,
				DeltaTx:     ip.DeltaTx,
				DNS:         dns,
				Dynamic:     ip.Dynamic,
				Gateway:     ip.Gateway,
				IP:          ip.IP,
				Mac:         ip.Mac,
				Name:        ip.Name,
				NetMask:     ip.NetMask,
				NetRx:       ip.NetRx,
				NetTx:       ip.NetTx,
				Rate:        ip.Rate,
			}
			NewIps = append(NewIps, netIp)
		}
	} else {
		for _, ip := range Ips {

			if ip.IP == "" {
				continue
			}

			var dns string
			if len(ip.DNS) > 0 {
				dns = ip.DNS[0]
			} else {
				dns = ""
			}

			netIp := response.NetIp{
				NetCardName: ip.NetCardName,
				Bandwidth:   ip.Bandwidth,
				DeltaRx:     ip.DeltaRx,
				DeltaTx:     ip.DeltaTx,
				DNS:         dns,
				Dynamic:     ip.Dynamic,
				Gateway:     ip.Gateway,
				IP:          ip.IP,
				Mac:         ip.Mac,
				Name:        ip.Name,
				NetMask:     ip.NetMask,
				NetRx:       ip.NetRx,
				NetTx:       ip.NetTx,
				Rate:        ip.Rate,
			}
			NewIps = append(NewIps, netIp)

		}
	}

	c.JSON(http.StatusOK, handle.Success(NewIps))
}

func (b *IpApi) IpSet(c *gin.Context) {

	req := response.IpSetting{}
	data, _ := io.ReadAll(c.Request.Body)
	_ = json.Unmarshal(data, &req)

	err := service.Valid(c.Request, req)
	if err != nil {
		logger.Error("error: ", err)
		errStr := fmt.Sprintf("%v", err)
		c.JSON(http.StatusUnprocessableEntity, handle.FailWithMsg(1, errStr))
		return
	}

	if global.DeviceType == "SE5" && (global.SdkVersion == "2.4.0" || global.SdkVersion == "2.7.0") {
		logger.Error("%s版本SDK不支持修改ip", global.SdkVersion)
		c.JSON(http.StatusUnprocessableEntity, handle.FailWithMsg(1, "不支持修改ip"))
		return
	}

	logger.Info("req:%v", req)

	var ip dto.IPSettings

	ip.Device = req.Device
	ip.Policy = Static.String()

	if req.IPType == Static.Code() {
		err = service.NotBlank(c.Request,
			"ip", req.IP,
			"subnetMask", req.SubnetMask)
		if err != nil {
			logger.Error("error: ", err)
			errStr := fmt.Sprintf("%v", err)
			c.JSON(http.StatusUnprocessableEntity, handle.FailWithMsg(1, errStr))
			return
		}
		ip.IP = req.IP
		ip.Mask = req.SubnetMask
		ip.Gateway = req.Gateway
		ip.DNS = req.DNS
	} else if req.IPType == Dynamic.Code() {
		ip.Policy = Dynamic.String()
	}

	logger.Info("ip:%v", ip)

	// 修改ip
	service.SaveOptLog(c.Request, "IP设置", "网卡"+req.Device+" 设置IP: "+req.IP)

	c.JSON(http.StatusOK, handle.OkWithMsg("请求成功"))

	go func() {
		time.Sleep(2 * time.Second)
		_, err = service.SetIP(ip)
		handle.HandleError(err, buserr.SetIpErr)
	}()

}

func (b *IpApi) GetTables(c *gin.Context) {
	sysTables, tables := getSsmTables()
	tablesList := IpTableLists{
		SysTables:  sysTables,
		UserTables: tables,
	}
	c.JSON(http.StatusOK, handle.Success(tablesList))

}

func (b *IpApi) AddTables(c *gin.Context) {
	req := dto.AddTable{}
	data, _ := io.ReadAll(c.Request.Body)
	_ = json.Unmarshal(data, &req)

	sysTables, tables := getSsmTables()
	for _, tab := range sysTables {
		if tab.SourcePort == req.SrcPort {
			logger.Error("端口 %s 已被映射\n", req.SrcPort)
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "端口已被映射"))
			return
		}
	}

	for _, tab := range tables {
		if tab.SourcePort == req.SrcPort {
			logger.Error("端口 %s 已被映射\n", req.SrcPort)
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "端口已被映射"))
			return
		}
	}

	// 查看端口是否已经被占用
	// 尝试监听指定端口
	if req.Protocol == "tcp" {
		ln, err := net.Listen(req.Protocol, "localhost:"+req.SrcPort)
		if err != nil {
			logger.Error("端口 %s 已被占用: %v\n", req.SrcPort, err)
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "端口已被占用"))
			return
		} else {
			// 如果端口没有被占用，关闭监听
			defer ln.Close()
			logger.Info("端口 %s 未被占用\n", req.SrcPort)
		}
	}
	req.Dirt = "in"
	req.Op = "append"

	res, _ := service.AddIpTable(req)
	if res.Code != 0 {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "add ip table failed"))
	}
	c.JSON(http.StatusOK, handle.Ok())
}

func (b *IpApi) DeleteTables(c *gin.Context) {
	req := struct {
		Num string `json:"num"`
	}{}
	data, _ := io.ReadAll(c.Request.Body)
	_ = json.Unmarshal(data, &req)

	res, _ := service.DeleteIpTable(req.Num)
	// fmt.Println(res)
	if res.Result == nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "delete ip table failed"))
		return
	}
	c.JSON(http.StatusOK, handle.Ok())
}

func getSsmTables() ([]Table, []Table) {
	list, err := service.GetIpTables()
	handle.HandleError(err)

	var sysTables []Table
	// fmt.Println(list)
	var i int
	for j, line := range list {
		if line == "" || j >= 26 {
			i = j
			break
		}
		if strings.HasPrefix(line, "Chain") || strings.HasPrefix(line, "num") {
			continue
		}
		var table Table
		var a string
		_, err := fmt.Sscanf(line, "%s %s %s %s %s %s %s %s %s", &table.Num, &table.Target, &table.Protocol, &a, &a, &table.SourceIP, &a, &table.SourcePort, &table.DestIp)
		if err != nil {
			logger.Error("解析行失败: %v\n", err)
			continue
		}

		if table.Target != "DNAT" {
			i = j + 1
			break
		}

		Index := strings.Index(table.SourcePort, ":")
		table.SourcePort = table.SourcePort[Index+1:]

		if strings.HasPrefix(table.DestIp, "to:") {
			parts := strings.Split(table.DestIp, ":")
			if len(parts) == 3 {
				table.DestIp = parts[1]
				table.DestPort = parts[2]
			}
		}

		// fmt.Println(table)
		sysTables = append(sysTables, table)
	}

	var tables []Table
	for i < len(list) {
		if list[i] == "" {
			break
		}

		var table Table
		var a string
		_, err := fmt.Sscanf(list[i], "%s %s %s %s %s %s %s %s %s", &table.Num, &table.Target, &table.Protocol, &a, &a, &table.SourceIP, &a, &table.SourcePort, &table.DestIp)
		if err != nil {
			logger.Error("解析行失败: %v\n", err)
			i++
			continue
		}
		if table.Target != "DNAT" {
			i = i + 1
			continue
		}

		Index := strings.Index(table.SourcePort, ":")
		table.SourcePort = table.SourcePort[Index+1:]

		if strings.HasPrefix(table.DestIp, "to:") {
			parts := strings.Split(table.DestIp, ":")
			if len(parts) == 3 {
				table.DestIp = parts[1]
				table.DestPort = parts[2]
			}
		}
		tables = append(tables, table)
		i++
	}
	return sysTables, tables
}

type Table struct {
	Num        string `json:"num"`
	Target     string `json:"target"`
	Protocol   string `json:"protocol"`
	SourceIP   string `json:"sourceIP"`
	SourcePort string `json:"sourcePort"`
	DestIp     string `json:"destIp"`
	DestPort   string `json:"destPort"`
}
type IpTableLists struct {
	SysTables  []Table `json:"sysTables"`
	UserTables []Table `json:"userTables"`
}

// 声明IP类型
type IpType int

const (
	None    IpType = iota
	Static         // 静态ip
	Dynamic        // 动态ip
)

func (c IpType) String() string {
	switch c {
	case None:
		return "None"
	case Static:
		return "static"
	case Dynamic:
		return "dhcp"
	}
	return "None"
}

func (c IpType) Code() int {
	switch c {
	case None:
		return -1
	case Static:
		return 1
	case Dynamic:
		return 2
	}
	return -1
}
