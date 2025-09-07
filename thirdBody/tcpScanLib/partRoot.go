package tcpScanLib

import (
	"net"
	"strings"
	"time"
)

type verifyFunc interface {
	Verifyer(conn net.Conn, port int) (bool, string)
}

type Http struct{}

func (root *Http) Verifyer(conn net.Conn, port int) (bool, string) {
	// HTTP服务验证
	req := "GET / HTTP/1.1\r\nHost: " + conn.RemoteAddr().String() + "\r\n\r\n"
	if _, err := conn.Write([]byte(req)); err != nil {
		return false, "HTTP探测失败"
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil || n == 0 {
		return false, "无HTTP响应"
	}

	// 检查是否包含HTTP响应特征
	response := string(buf[:n])
	if len(response) >= 4 && response[:4] == "HTTP" {
		return true, "HTTP"
	}
	return false, "非HTTP服务"
}

type Https struct{}

func (root *Https) Verifyer(conn net.Conn, port int) (bool, string) {
	// HTTPS服务验证（简单检查TLS握手）
	// 发送不完整的ClientHello，看是否有响应
	tlsProbe := []byte{0x16, 0x03, 0x01, 0x00, 0x01} // TLS ClientHello开始
	if _, err := conn.Write(tlsProbe); err != nil {
		return false, "TLS探测失败"
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil || n == 0 {
		return false, "无TLS响应"
	}
	return true, "HTTPS"
}

type MySQL struct{}

func (root *MySQL) Verifyer(conn net.Conn, port int) (bool, string) {
	// MySQL服务验证
	// 发送MySQL客户端握手初始化包
	mysqlProbe := []byte{0x05, 0x00, 0x00, 0x01, 0x85, 0xa6, 0x3f, 0x00}
	if _, err := conn.Write(mysqlProbe); err != nil {
		return false, "MySQL探测失败"
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil || n == 0 {
		return false, "无MySQL响应"
	}
	return true, "MySQL"
}

type SSH struct{}

func (root *SSH) Verifyer(conn net.Conn, port int) (bool, string) {
	// SSH服务验证
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil || n == 0 {
		return false, "无SSH响应"
	}

	if len(buf) >= 4 && strings.Contains(string(buf), "SSH") {
		return true, "SSH"
	}
	return false, "非SSH服务"
}

type DefaultService struct{}

func (root *DefaultService) Verifyer(conn net.Conn, port int) (bool, string) {
	// 通用验证：尝试读取数据或发送简单探测
	// 首先尝试读取可能的服务Banner

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err == nil && n > 0 {
		return true, "未知服务 (有响应)"
	}

	// 如果没有数据，尝试发送简单探测包
	probe := []byte("HELLO\r\n")
	if _, err := conn.Write(probe); err != nil {
		return false, "通用探测失败"
	}

	// 再次尝试读取响应
	n, err = conn.Read(buf)
	if err == nil && n > 0 {
		return true, "未知服务 (有响应)"
	}

	// 对于没有响应的端口，尝试保持连接一小段时间
	time.Sleep(800 * time.Millisecond)
	// 检查连接是否仍然有效
	if _, err := conn.Write([]byte(" ")); err == nil {
		return false, "未知服务 (连接有效)"
	}

	return false, "虚假开放端口"
}
