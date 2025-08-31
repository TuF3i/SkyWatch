package GoPing

import (
	"fmt"
	"testing"
)

// main函数是程序的入口点
func TestRun(t *testing.T) {
	// 初始化一个Ping实例，用于对指定的域名进行ping操作
	// 参数"baidu.com"是目标域名，8是ping的次数，Data是ping的数据包内容
	ping, err := New("172.20.12.141", 2)
	// 如果初始化过程中出现错误，则输出错误信息并终止程序
	if err != nil {
		fmt.Println(err)
	}
	// 执行ping操作，这里指定ping的次数为5
	err = ping.Ping(5)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 输出ping操作的总次数，这里查询的是次数6的ping操作次数
	fmt.Println(ping.PingCount(6))
	err = ping.Close()
	if err != nil {
		t.Error(err)
		return
	}
}
