package GoPing

import (
	"errors"
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"net"
	"os"
	"time"
)

// MarshalMsg 根据请求编号和数据生成ICMP回显请求消息。
// req: 请求编号，用于标识特定的请求。
// data: 需要包含在ICMP消息中的数据。
// 返回值: 编码后的ICMP消息字节流以及可能的错误。
func MarshalMsg(req int, data []byte) ([]byte, error) {
	// 获取当前进程ID并将其低16位用作ICMP消息的ID。
	// 这有助于唯一标识发送的ICMP请求。
	xid, seq := os.Getpid()&0xffff, req

	// 构造ICMP消息体。
	// 使用ICMP回显请求类型和代码0。
	// 设置ID和序列号为之前计算的值，并包含传入的数据。
	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: xid, Seq: seq,
			Data: data,
		},
	}

	// 将构造好的ICMP消息体编码为字节流。
	// 如果编码过程中发生错误，将返回该错误。
	return wm.Marshal(nil)
}

// sendPingMsg 发送一个ICMP ping请求并接收回复。
// 参数c是网络连接，wb是写入缓冲区的字节数据。
// 返回值reply包含回复的信息，如延迟时间、TTL和错误信息。
func sendPingMsg(c net.Conn, wb []byte) (reply Reply) {
	// 记录发送请求的时间
	start := time.Now()

	// 尝试向连接c写入wb数据，如果写入失败，则将错误信息存储在reply中并返回
	if _, reply.Error = c.Write(wb); reply.Error != nil {
		return
	}

	// 准备一个缓冲区以接收回复数据
	rb := make([]byte, 1500)
	var n int
	// 尝试从连接c读取数据，如果读取失败，则将错误信息存储在reply中并返回
	n, reply.Error = c.Read(rb)
	if reply.Error != nil {
		return
	}

	// 计算发送请求到接收回复的时间间隔
	duration := time.Now().Sub(start)
	// 从接收的數據中提取TTL值
	ttl := rb[8]
	// 处理接收的數據，去除ICMP头部，保留数据部分
	rb = func(b []byte) []byte {
		if len(b) < 20 {
			return b
		}
		hdrlen := int(b[0]&0x0f) << 2
		return b[hdrlen:]
	}(rb)
	// 尝试解析接收的ICMP消息，如果解析失败，则将错误信息存储在reply中并返回
	var rm *icmp.Message
	rm, reply.Error = icmp.ParseMessage(1, rb[:n])
	if reply.Error != nil {
		return
	}

	// 根据解析出的ICMP消息类型，设置reply的值
	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		// 如果是回显回复，则计算并设置延迟时间，TTL，并清空错误信息
		t := int64(duration / time.Millisecond)
		reply = Reply{t, ttl, nil}
	case ipv4.ICMPTypeDestinationUnreachable:
		// 如果是目标不可达，则设置相应的错误信息
		reply.Error = errors.New("destination Unreachable")
	default:
		// 对于其他类型的消息，设置相应的错误信息
		reply.Error = fmt.Errorf("not ICMPTypeEchoReply %v", rm)
	}
	return
}
