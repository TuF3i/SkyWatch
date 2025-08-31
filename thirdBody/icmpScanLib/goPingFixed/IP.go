package GoPing

import (
	"gitee.com/liumou_site/logger"
	"net"
)

// IsIP 检查给定的字符串是否为有效的IP地址。
// 它首先尝试将字符串解析为一个IPv4或IPv6地址。如果解析失败，
// 则尝试将字符串解析为一个主机名，并通过DNS查找验证该主机名是否存在。
// 参数:
//
//	ip - 待检查的字符串，可能是IP地址或主机名。
//
// 返回值:
//
//	如果字符串是有效的IP地址或可解析的主机名，则返回true；否则返回false。
func IsIP(ip string) bool {
	// 尝试将字符串解析为IP地址
	// 判断ip地址是否合法
	if net.ParseIP(ip) == nil {
		// 如果解析失败，尝试将字符串解析为主机名
		// 判断是否为域名
		if _, err := net.LookupHost(ip); err != nil {
			// 如果主机名查找也失败，则返回false
			return false
		} else {
			// 如果主机名查找成功，则返回true
			return true
		}
	}
	// 如果字符串成功解析为IP地址，则直接返回true
	return true
}

// FilterIP 过滤给定的IP地址列表，返回有效的IP地址数组。
// 该函数遍历每个输入的IP地址字符串，检查其是否为有效的IP地址。
// 如果是有效的IP地址，则将其添加到结果数组中；如果不是，则记录错误。
// 参数:
//
//	ips []string: 待检查的IP地址列表。
//
// 返回值:
//
//	[]string: 有效的IP地址列表。
func FilterIP(ips []string) []string {
	var validIps []string

	// 遍历输入的IP地址列表
	for _, ip := range ips {
		// 检查当前IP地址是否有效
		if IsIP(ip) {
			// 如果有效，则将其添加到结果列表中
			validIps = append(validIps, ip)
		} else {
			// 如果无效，则记录错误日志
			logger.Error("Not an IP address or domain name: ", ip)
		}
	}

	// 返回有效的IP地址列表
	return validIps
}
