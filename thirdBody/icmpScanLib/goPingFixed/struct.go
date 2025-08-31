package GoPing

import (
	"net"
	"sync"
	"time"
)

// PingSet 定义了一个用于网络ping操作的集合，包含所有必要的参数和状态。
// 它旨在支持自定义地址、数据包内容和超时设置的ping操作。
type PingSet struct {
	Addr    string   // 地址指定要ping的目标主机或网络设备的IP地址或域名。
	Conn    net.Conn // Conn 保存了与目标主机建立的网络连接，用于发送和接收数据包。
	Data    []byte   // Data 包含了要发送的数据包内容，通常为"ICMP ECHO REQUEST"的数据部分。
	Timeout time.Duration      // Timeout 指定了操作的超时时间，以秒为单位，用于控制ping操作的最长等待时间。
	Req     int      // Req 表示要发送的ICMP ECHO REQUEST的数量,默认8
	Print   bool     // Print 是否打印ping明细信息,默认：true
}

// Reply 结构体用于封装回复信息。
// 它包含回复的时间戳、生存时间（TTL）和可能的错误信息。
type Reply struct {
	Time  int64 // 时间戳表示回复的生成时间。
	TTL   uint8 // TTL 表示回复消息的生存时间，用于控制消息的有效性。
	Error error // Error 保存操作过程中可能发生的错误信息。
}

// Concurrency 创建并发结构
type Concurrency struct {
	Addr   []string // 需要并发ping的地址清单
	Res    sync.Map // 结果
	Result map[string]bool
	Ch     chan string // 结果通道
	L      sync.Mutex  // 增加锁
	// 统计结果
	Total   int32 // 总数
	Success int32 // 成功
	Fail    int32 // 失败
	Err     error // 错误
}
