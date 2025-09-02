package tcpScanLib

import (
	"errors"
	"net"
	"strconv"
	"time"
)

// TCPPortScan 检查指定IP的TCP端口是否开放
// 参数:
//
//	ip: 目标IP地址
//	port: 目标端口号
//	timeout: 连接超时时间
//
// 返回值:
//
//	bool: 端口是否开放
//	error: 可能的错误信息，如无效的IP或端口
func TCPPortScan(ip string, port int, timeout time.Duration) (bool, error) {
	// 验证端口是否在有效范围内
	if port < 1 || port > 65535 {
		return false, errors.New("无效的端口号，必须在1-65535之间")
	}

	// 验证IP地址格式是否有效
	if net.ParseIP(ip) == nil {
		return false, errors.New("无效的IP地址格式")
	}

	// 构建地址字符串（IP:端口）
	address := net.JoinHostPort(ip, strconv.Itoa(port))

	// 尝试建立TCP连接，设置超时时间
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		// 检查是否是超时错误
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			// 超时错误，端口可能被过滤或防火墙阻止，不视为错误但端口不开放
			return false, nil
		}
		// 其他错误，如连接被拒绝（端口关闭）等
		return false, err
	}
	defer conn.Close() // 确保连接被关闭

	// 连接成功，端口开放
	return true, nil
}
