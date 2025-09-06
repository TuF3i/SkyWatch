package tcpScanLib

import (
	logger "SkyWatch/thirdBody/gologger"
	"errors"
	"io"
	"net"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// TCPPortScan 检查指定IP的TCP端口状态，增加数据返回检测
// 参数:
//
//	ip: 目标IP地址
//	port: 目标端口号
//	timeout: 连接超时时间
//
// 返回值:
//
//	bool: 端口是否开放(只有open返回true，其他状态返回false)
//	error: 可能的错误信息，如无效的IP或端口或网络错误
func TCPPortScan(ip string, port int, timeout time.Duration) (bool, error) {
	logs := logger.NewLogger(1)
	logs.Modular = "tcpScanLib"

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

	// 使用更精确的 Dialer
	dialer := &net.Dialer{
		Timeout: timeout,
		Control: func(network, address string, c syscall.RawConn) error {
			// 这里可以添加底层socket控制，但当前不需要
			return nil
		},
	}

	// 尝试建立TCP连接
	conn, err := dialer.Dial("tcp", address)
	if err != nil {
		// 获取错误字符串以便检查
		errStr := err.Error()

		// 检查超时错误 (filtered)
		if isTimeoutError(err) {
			return false, nil
		}

		// 检查连接被拒绝错误 (closed)
		if isConnectionRefused(err) {
			return false, nil
		}

		// 检查连接重置错误 (reset)
		if isConnectionReset(err) {
			return false, nil
		}

		// 检查网络不可达错误
		if isNoRouteToHost(err) {
			return false, errors.New("网络不可达: " + errStr)
		}

		// 其他网络错误
		return false, errors.New("网络错误: " + errStr)
	}
	defer conn.Close() // 确保连接被关闭

	// 关键改进：设置读取超时，检测端口是否返回数据
	readTimeout := 1 * time.Second // 数据读取超时时间
	if err := conn.SetReadDeadline(time.Now().Add(readTimeout)); err != nil {
		logger.Debug("设置读取超时失败: %v", err)
		return false, nil
	}

	// 尝试读取少量数据（1字节）
	buf := make([]byte, 1)
	n, err := conn.Read(buf)
	if err != nil {
		// 处理读取错误
		if errors.Is(err, io.EOF) {
			// 连接建立后立即关闭，无数据返回
			logger.Debug("端口 %d 连接建立但无数据返回（EOF）", port)
			return false, nil
		} else if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			// 读取超时，无数据返回
			logger.Debug("端口 %d 连接建立但超时无数据返回", port)
			return false, nil
		} else {
			// 其他读取错误
			logger.Debug("端口 %d 数据读取错误: %v", port, err)
			return false, nil
		}
	}

	// 成功读取到数据，确认端口开放
	if n > 0 {
		logger.Info("[+] %v:%v is open! 检测到返回数据", ip, port)
		return true, nil
	}

	// 未读取到数据且无错误（理论上不会走到这里）
	logger.Debug("端口 %d 连接建立但未返回任何数据", port)
	return false, nil
}

// isTimeoutError 检查错误是否是超时错误
func isTimeoutError(err error) bool {
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		return true
	}

	// 检查错误字符串中的超时关键词
	errStr := err.Error()
	return strings.Contains(errStr, "timeout") ||
		strings.Contains(errStr, "timed out") ||
		strings.Contains(errStr, "i/o timeout")
}

// isConnectionRefused 检查错误是否是连接拒绝错误
func isConnectionRefused(err error) bool {
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		if opErr.Op == "dial" {
			// 在不同操作系统上，连接拒绝的错误消息可能不同
			errStr := opErr.Err.Error()
			return strings.Contains(errStr, "connection refused") ||
				strings.Contains(errStr, "actively refused") ||
				strings.Contains(errStr, "refused") ||
				strings.Contains(errStr, "目标计算机积极拒绝") // 中文Windows的错误消息
		}
	}
	return false
}

// isConnectionReset 检查错误是否是连接重置错误
func isConnectionReset(err error) bool {
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		errStr := opErr.Err.Error()
		return strings.Contains(errStr, "connection reset") ||
			strings.Contains(errStr, "reset by peer")
	}
	return false
}

// isNoRouteToHost 检查错误是否是 "no route to host"
func isNoRouteToHost(err error) bool {
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		errStr := opErr.Err.Error()
		return strings.Contains(errStr, "no route to host") ||
			strings.Contains(errStr, "host is down") ||
			strings.Contains(errStr, "network is unreachable")
	}
	return false
}
