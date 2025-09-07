package tcpScanLib

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	logger "SkyWatch/thirdBody/gologger"
)

var logs *logger.LocalLogger = logger.NewLogger(1)

// 扫描配置常量，可根据网络环境调整
const (
	defaultTimeout = 3 * time.Second        // 默认超时时间
	retryMax       = 2                      // 最大重试次数
	retryDelay     = 500 * time.Millisecond // 重试间隔
	minPort        = 1
	maxPort        = 65535
)

// TCPPortScan 检查指定IP的TCP端口状态（高可靠性版本）
// 增加重试机制和智能超时策略，降低丢包率
func TCPPortScan(ip string, port int, timeout time.Duration) (bool, error) {
	logs.Modular = "tcpScanLib"

	// 1. 参数验证
	if err := validateParams(ip, port); err != nil {
		return false, err
	}

	// 2. 设置超时（确保有合理的最小值）
	effectiveTimeout := timeout
	if effectiveTimeout < time.Second {
		effectiveTimeout = defaultTimeout
	}

	address := net.JoinHostPort(ip, strconv.Itoa(port))
	lastErr := error(nil)

	// 3. 带重试的连接尝试
	for attempt := 0; attempt <= retryMax; attempt++ {
		// 重试时增加超时时间（指数退避）
		attemptTimeout := effectiveTimeout + time.Duration(attempt)*effectiveTimeout/2
		conn, err := net.DialTimeout("tcp", address, attemptTimeout)

		if err == nil {
			_ = conn.Close() // 确保连接关闭
			logs.Info("[+] %s:%d is open", ip, port)
			return true, nil
		}

		// 4. 错误分类处理
		switch {
		case isTimeoutError(err):
			// 超时错误可能是临时网络波动，继续重试
			lastErr = fmt.Errorf("连接超时 (尝试 %d/%d)", attempt+1, retryMax+1)
			//logs.Debug("端口 %d 超时，将重试 (尝试 %d/%d)", port, attempt+1, retryMax+1)

		case isConnectionRefused(err), isConnectionReset(err):
			// 明确的关闭状态，无需重试
			//logs.Debug("端口 %d 关闭 (尝试 %d/%d)", port, attempt+1, retryMax+1)
			return false, nil

		case isNoRouteToHost(err):
			// 网络不可达，终止重试
			return false, fmt.Errorf("网络不可达: %w", err)

		default:
			// 其他错误，根据情况决定是否重试
			lastErr = fmt.Errorf("扫描错误: %w", err)
			//logs.Debug("端口 %d 扫描错误: %v (尝试 %d/%d)", port, err, attempt+1, retryMax+1)
		}

		// 5. 最后一次尝试不延迟
		if attempt < retryMax {
			time.Sleep(retryDelay)
		}
	}

	// 所有重试都失败
	return false, lastErr
}

// 参数验证
func validateParams(ip string, port int) error {
	if port < minPort || port > maxPort {
		return fmt.Errorf("无效的端口号 %d，必须在%d-%d之间", port, minPort, maxPort)
	}

	if net.ParseIP(ip) == nil {
		return fmt.Errorf("无效的IP地址格式: %s", ip)
	}

	return nil
}

// isTimeoutError 增强版超时错误判断
func isTimeoutError(err error) bool {
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return true
	}

	errStr := strings.ToLower(err.Error())
	return strings.Contains(errStr, "timeout") ||
		strings.Contains(errStr, "timed out") ||
		strings.Contains(errStr, "i/o timeout") ||
		strings.Contains(errStr, "连接超时")
}

// isConnectionRefused 增强版连接拒绝错误判断
func isConnectionRefused(err error) bool {
	errStr := strings.ToLower(err.Error())
	return strings.Contains(errStr, "connection refused") ||
		strings.Contains(errStr, "actively refused") ||
		strings.Contains(errStr, "refused") ||
		strings.Contains(errStr, "目标计算机积极拒绝") ||
		strings.Contains(errStr, "拒绝连接")
}

// isConnectionReset 增强版连接重置错误判断
func isConnectionReset(err error) bool {
	errStr := strings.ToLower(err.Error())
	return strings.Contains(errStr, "connection reset") ||
		strings.Contains(errStr, "reset by peer") ||
		strings.Contains(errStr, "wsarecv") ||
		strings.Contains(errStr, "连接被重置")
}

// isNoRouteToHost 增强版网络不可达错误判断
func isNoRouteToHost(err error) bool {
	errStr := strings.ToLower(err.Error())
	return strings.Contains(errStr, "no route to host") ||
		strings.Contains(errStr, "host is down") ||
		strings.Contains(errStr, "network is unreachable") ||
		strings.Contains(errStr, "找不到主机") ||
		strings.Contains(errStr, "网络不可达")
}
