package network

import (
	"fmt"
	"github.com/miaoyc666/goArsenal/file"
	"github.com/miaoyc666/goArsenal/system"
	"net"
	"strings"
)

/*
File name    : network.go
Author       : miaoyc
Create date  : 2021/12/2 2:24 下午
Description  : 网卡网络相关
*/


//GetAllNetworkDeviceAddrs 获取所有网卡的设备地址
func GetAllNetworkDeviceAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return macAddrs
	}
	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}

//GetVirtualNetworkDeviceNames 获取虚拟网卡名称列表
func GetVirtualNetworkDeviceNames() ([]string, error) {
	output, err := system.RunSystemShell("/bin/bash", "-c", "ls /sys/devices/virtual/net")
	if err != nil {
		return nil, err
	}
	output = strings.TrimSpace(output)
	blackList := strings.Split(output, "\n")
	return blackList, nil
}

//GetPhysicalNetworkDeviceNames 获取物理网卡名称列表, 相关命令：ls /sys/class/net/ | grep -v "$(ls /sys/devices/virtual/net/)"
func GetPhysicalNetworkDeviceNames() ([]string, error) {
	output, err := system.RunSystemShell("/bin/bash", "-c", "ls /sys/class/net/ | grep -v \"$(ls /sys/devices/virtual/net/)\"")
	if err != nil {
		return nil, err
	}
	output = strings.TrimSpace(output)
	blackList := strings.Split(output, "\n")
	return blackList, nil
}

//GetPhysicalNetworkDeviceAddrs 获取物理网卡设备地址, 使用lspci来获取网卡设备ID目的是保证获取到的顺序是网卡设备的实际插卡顺序
func GetPhysicalNetworkDeviceAddrs() (macAddrs []string) {
	netInterfaces, err := getIfIdList()
	if err != nil {
		return macAddrs
	}
	for _, netInterface := range netInterfaces {
		mac, _ := getMacById(netInterface)
		macAddrs = append(macAddrs, mac)
	}
	return macAddrs
}

//getIfIdList 获取网卡在pci总线上的设备ID
func getIfIdList() ([]string, error) {
	output, err := system.RunSystemShell("/bin/bash", "-c", "lspci -D |grep Eth |awk '{print $1}'")
	if err != nil {
		return nil, err
	}
	output = strings.TrimSpace(output)
	ifIds := strings.Split(output, "\n")
	return ifIds, nil
}

// getMacById 通过设备ID获取mac地址
/*
	业务流程：
	1.首先判断服务器是物理机还是虚拟机，物理机会存在/sys/bus/pci/devices/{ifId}/net目录, virtio虚拟机不存在该目录
	2.根据服务器类型，取不同位置的address文件中的mac地址
 */
func getMacById(ifId string) (string, error) {
	dirPath := fmt.Sprintf("/sys/bus/pci/devices/%s/net", ifId)
	queryMac := ""
	if file.IsDir(dirPath) {
		queryMac = fmt.Sprintf("cat /sys/bus/pci/devices/%s/net/*/address", ifId)
	} else {
		queryMac = fmt.Sprintf("cat /sys/bus/pci/devices/%s/virtio0/net/*/address", ifId)
	}
	output, err := system.RunSystemShell("/bin/bash", "-c", queryMac)
	if err != nil {
		return "", err
	}
	output = strings.TrimSpace(output)
	macList := strings.Split(output, "\n")
	return macList[0], nil
}
