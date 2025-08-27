package ipProcesser

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func IpGenerater(ipRange string) ([]string, error) {
	// 检查是否是CIDR格式
	if strings.Contains(ipRange, "/") {
		return parseCIDR(ipRange)
	}

	// 检查是否是IP范围 (如 192.168.1.1-100)
	if strings.Contains(ipRange, "-") {
		return parseIPSequence(ipRange)
	}

	// 单个IP
	return []string{ipRange}, nil
}

// 解析CIDR格式的IP范围
func parseCIDR(cidr string) ([]string, error) {
	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ipnet.IP.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	// 移除网络地址和广播地址
	if len(ips) > 2 {
		return ips[1 : len(ips)-1], nil
	}
	return ips, nil
}

// IP地址递增
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// 解析IP序列 (如 192.168.1.1-100)
func parseIPSequence(ipSeq string) ([]string, error) {
	parts := strings.Split(ipSeq, ".")
	if len(parts) != 4 {
		return nil, fmt.Errorf("无效的IP范围格式")
	}

	lastPart := parts[3]
	if !strings.Contains(lastPart, "-") {
		return []string{ipSeq}, nil
	}

	rangeParts := strings.Split(lastPart, "-")
	if len(rangeParts) != 2 {
		return nil, fmt.Errorf("无效的IP范围格式")
	}

	start, err := strconv.Atoi(rangeParts[0])
	if err != nil {
		return nil, fmt.Errorf("无效的起始IP")
	}

	end, err := strconv.Atoi(rangeParts[1])
	if err != nil {
		return nil, fmt.Errorf("无效的结束IP")
	}

	if start > end {
		return nil, fmt.Errorf("起始IP不能大于结束IP")
	}

	var ips []string
	base := strings.Join(parts[0:3], ".")
	for i := start; i <= end; i++ {
		ips = append(ips, fmt.Sprintf("%s.%d", base, i))
	}

	return ips, nil
}

func ReadIPsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(fmt.Sprintf("无法打开文件: %v", err))
		return nil, fmt.Errorf("无法打开文件: %v", err)
	}
	defer file.Close()

	var ips []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		ips = append(ips, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(fmt.Sprintf("读取文件时出错: %v", err))
		return nil, fmt.Errorf("读取文件时出错: %v", err)
	}

	return ips, nil
}
