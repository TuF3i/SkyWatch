package icmpScanLib

import (
	"SkyWatch/thirdBody/icmpScanLib/goPingFixed"
	"fmt"
	"time"
)

func IsHostAlive(ipaddr string) (bool, error) {
	ping, err := GoPing.New(ipaddr, 3*time.Second)
	defer ping.Close()
	if err != nil {
		return false, fmt.Errorf("创建ping实例失败:", err)
	}
	err = ping.Ping(3)
	if err != nil {
		return false, fmt.Errorf("ping测试失败", err)
	} else {
		return true, nil
	}
}
