package network

// 获取本地相关的网络信息

import (
	"net"
	"regexp"
	"strconv"
)

// 网卡/网络适配器信息结构体
type Adaptor struct {
	// 网卡序号
	Index int
	// 最大传输单元
	Mtu int
	// 网卡名称
	Name string
	// MAC地址
	Mac net.HardwareAddr
	// 状态标识
	Flags net.Flags
	// IPv4
	IPv4 net.IP
	// IPv4子网掩码个数
	IPv4MaskCount int
	// IPv6
	IPv6 net.IP
	// IPv6子网掩码个数
	IPv6MaskCount int
}

// 1.获取所有有效的网卡信息
func Adaptors() ([]*Adaptor, error) {
	// 获取所有的网卡
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	adaptors := make([]*Adaptor, 0)
	// 遍历所有的网卡
LAB:
	for _, iface := range ifaces {
		// 网卡过滤 : 网卡状态Flags为up
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		// ip地址过滤 :
		// 1.ipv4和ipv6同时存在 addresses.len > 1 ,ipv4,ipv6转换不是nil，ipv4和ipv6的掩码个数为正数
		// 2.非本地回环
		addresses, err := iface.Addrs()
		if err != nil || len(addresses) < 2 {
			continue
		}
		var (
			ipv4, ipv6                   net.IP
			ipv4MaskCount, ipv6MaskCount int
		)
		for _, address := range addresses {
			if ipnet, ok := address.(*net.IPNet); ok && ipnet.IP.IsLoopback() {
				continue LAB
			} else {
				if ipnet.IP.To4() != nil {
					ipv4 = ipnet.IP.To4()
					cs := regexp.MustCompile(`/\d+`).FindString(address.String())
					cs = regexp.MustCompile(`/`).ReplaceAllString(cs, "")
					c, err := strconv.Atoi(cs)
					if err != nil {
						ipv4MaskCount = -1
						continue LAB
					}
					ipv4MaskCount = c
				} else if ipnet.IP.To16() != nil {
					ipv6 = ipnet.IP.To16()
					cs := regexp.MustCompile(`/\d+`).FindString(address.String())
					cs = regexp.MustCompile(`/`).ReplaceAllString(cs, "")
					c, err := strconv.Atoi(cs)
					if err != nil {
						ipv6MaskCount = -1
						continue LAB
					}
					ipv6MaskCount = c
				}
			}
		}
		if ipv4 == nil || ipv6 == nil || ipv4MaskCount < 0 || ipv6MaskCount < 0 {
			continue
		}
		// mac地址判断 : 不允许为nil
		if iface.HardwareAddr == nil {
			continue
		}
		adaptor := &Adaptor{
			Index:         iface.Index,
			Mtu:           iface.MTU,
			Name:          iface.Name,
			Mac:           iface.HardwareAddr,
			Flags:         iface.Flags,
			IPv4:          ipv4,
			IPv4MaskCount: ipv4MaskCount,
			IPv6:          ipv6,
			IPv6MaskCount: ipv6MaskCount,
		}
		adaptors = append(adaptors, adaptor)
	}
	return adaptors, nil
}
