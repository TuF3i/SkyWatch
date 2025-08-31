package GoPing

import (
	"fmt"
	"gitee.com/liumou_site/gns"
	"gitee.com/liumou_site/logger"
	"testing"
)

// TestConcurrency_Start 测试并发启动功能。
// 本函数的目的是验证并发处理IP列表并标记它们的在线/离线状态的功能是否有效。
func TestConcurrency_Start(t *testing.T) {
	var ips []string

	// 初始化一个IP子网
	// 获取本机IP网段
	sub := gns.NewIp()
	// 获取一个可用的IP地址
	ip, err := sub.GetUseIP()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("本机IP网段：", gns.IpCutSubnet(ip))
	logger.Info("本机IP地址：", ip)
	// 使用for循环生成IP列表
	for i := 0; i < 255; i++ {
		ip_ := fmt.Sprintf("%s.%d", gns.IpCutSubnet(ip), i)
		logger.Debug("IP地址：", ip_)
		ips = append(ips, ip_)
	}
	// 创建一个并发处理实例
	p := NewConcurrency(ips)
	// 启动并发检测
	p.Start()
	// 遍历结果，输出每个IP的在线/离线状态
	for k, v := range p.Result {
		if v {
			logger.Info(k, " : 在线")
		} else {
			logger.Info(k, " : 离线")
		}
	}
	// 输出成功的数量、失败的数量以及总共处理的IP数量
	logger.Info("成功：", p.Success, " 失败：", p.Fail, " 总数：", p.Total)
}
