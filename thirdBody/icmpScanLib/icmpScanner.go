package icmpScanLib

import (
	"gitee.com/liumou_site/logger"

	"SkyWatch/thirdBody/icmpScanLib/goPingFixed"

	"fmt"
	"time"
)

var _PING_COUNT int = 1
var _PING_SIZE int = 1400
var _PING_TIMEOUT time.Duration = time.Second * 5

func IsHostAlive(ipaddr string) (bool, error) {

	//ping.Logger.Infof("\nping:%s\n", ipaddr)
	logs := logger.NewLogger(1)
	logs.Modular = "icmpScanLib"
	pinger, err := ping.NewPinger(ipaddr)
	if err != nil {
		logs.Error("创建ping实例失败:", err)
		return false, fmt.Errorf("创建ping实例失败:", err)
	}
	pinger.Count = _PING_COUNT
	pinger.Size = _PING_SIZE
	pinger.Timeout = _PING_TIMEOUT
	pinger.SetPrivileged(true)

	err = pinger.Run()

	if err != nil {
		logs.Error("ping异常：%s\n", err.Error())
		//fmt.Printf("ping异常：%s\n", err.Error())
		return false, fmt.Errorf("ping异常：%s\n", err.Error())
	}
	stats := pinger.Statistics()
	// 如果回包大于等于1则判为ping通
	if stats.PacketsRecv >= 1 {
		logger.Debug("[+]IPaddr %v is alive!", ipaddr)
		return true, nil
	} else {
		//fmt.Printf("IP can not reach:%s\n", ipaddr)
		return false, fmt.Errorf("IP can not reach:%s\n", ipaddr)
	}

}
