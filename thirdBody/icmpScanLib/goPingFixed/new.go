package GoPing

import (
	"gitee.com/liumou_site/logger"
	"sync"
)

// NewConcurrency 初始化并返回一个 Concurrency 实例，该实例用于并发地对多个IP地址进行ping操作。
// ips: 待检查的IP地址列表。
// 返回值: 一个 Concurrency 实例，包含了IP地址列表、结果通道、等待组等用于并发操作的字段。
func NewConcurrency(ips []string) *Concurrency {
	// 检查输入的IP地址列表是否为空
	if len(ips) == 0 {
		logger.Error("ips is nil")
		// 如果为空，则返回一个空的Concurrency实例
		return &Concurrency{}
	}

	// 过滤ips列表，剔除非IP地址的字符串
	// 剔除非ip
	ips = FilterIP(ips)

	// 创建一个字符串通道，用于接收ping操作的结果
	// 初始化
	// 创建一个缓冲通道，用于接收ping的结果
	ch := make(chan string, len(ips))
	// 创建锁
	l := sync.Mutex{}
	// 返回一个初始化好的Concurrency实例
	return &Concurrency{
		Addr:    ips,
		Res:     sync.Map{},
		Result:  make(map[string]bool),
		Ch:      ch,
		L:       l,
		Total:   0,
		Success: 0,
		Fail:    0,
		Err:     nil,
	}
}
