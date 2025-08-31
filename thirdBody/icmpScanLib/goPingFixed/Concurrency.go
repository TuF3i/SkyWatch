package GoPing

import (
	"gitee.com/liumou_site/logger"
	"sync"
	"sync/atomic"
)

// Start 启动并发任务来检测一组地址的连通性。
// 该方法首先为每个地址添加到等待组中，然后并发地执行ping操作，最后等待所有ping操作完成并处理结果。
func (c *Concurrency) Start() {
	// 为等待组添加与地址列表长度相等的数量，确保所有地址的检测都已被计入等待中。
	// 打印有效的IP
	var wg sync.WaitGroup
	// 遍历地址列表，对每个地址启动一个goroutine进行ping检测。
	for _, v := range c.Addr {
		wg.Add(1)
		go c.ping(v, &wg)
	}
	// 等待所有goroutine完成，确保所有地址的ping检测都已完成。
	wg.Wait()
	// 所有ping检测完成后，处理检测结果，这可能包括对结果的分析或报告。
	c.handel()
}

// handel 处理并发情况下的IP地址检查。
// 该方法遍历IP地址列表，并通过原子操作更新成功和失败的计数器，以及总处理数。
func (c *Concurrency) handel() {
	// 遍历IP地址列表
	for _, ip_ := range c.Addr {
		// 使用Load方法安全地检查IP地址是否已在结果映射中
		if _, ok := c.Res.Load(ip_); ok {
			// 如果IP地址存在于结果映射中，表示处理成功
			c.Result[ip_] = true
			atomic.AddInt32(&c.Success, 1)
		} else {
			// 如果IP地址不存在于结果映射中，表示处理失败
			c.Result[ip_] = false
			atomic.AddInt32(&c.Fail, 1)
		}
		// 更新总处理数
		atomic.AddInt32(&c.Total, 1)
	}
}

// ping 使用并发方式对给定的IP地址执行ping操作。
// 该方法是Concurrency类型的一个方法，它通过一个等待组(Wg)来协调并发操作的结束。
// 参数:
//
//	ip - 需要进行ping操作的IP地址。
func (c *Concurrency) ping(ip string, wg *sync.WaitGroup) {
	// 使用defer确保在方法返回前调用Wg.Done()，表示该并发任务已完成。
	defer wg.Done()
	// 记录ping操作开始的信息。
	logger.Info("ping start:", ip)

	// 尝试创建一个新的Ping实例。
	p, err := New(ip, 5)
	p.Print = false
	defer func(p *PingSet) {
		err := p.Close()
		if err != nil {
			logger.Error(err)
		}
	}(p)
	//如果创建失败，记录错误并返回。
	if err != nil {
		logger.Error("new error:", err)
		return
	}

	// 对IP执行ping操作。
	err = p.Ping(5)
	// 如果ping操作失败，记录错误。
	if err != nil {
		logger.Error("ping error:", err)
		//c.Res.Store(ip, "false")
	} else {
		// 如果ping操作成功，记录成功信息并标记该IP为可达。
		logger.Info("ping success:", ip)
		c.L.Lock()
		defer c.L.Unlock()
		c.Res.Store(ip, "true")
	}
}
