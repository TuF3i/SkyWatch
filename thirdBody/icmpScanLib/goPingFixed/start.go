package GoPing

import (

	"gitee.com/liumou_site/logger"
	"net"
	"time"
	"fmt"
)

// Ping 方法用于向远程主机发送ping请求。
// 参数 count 指定发送ping请求的次数。
func (pg *PingSet) Ping(count int) error {
	//defer func(pg *PingSet) {
	//	err := pg.Close()
	//	if err != nil {
	//		logger.Error(err)
	//	}
	//}(pg)
	// 尝试连接远程主机，如果失败则打印错误信息并返回。
	if err := pg.Dail(); err != nil {
		logger.Error("Not found remote host")
		return err
	}

	// 打印ping操作的起始地址。
	if pg.Print {
		//fmt.Println("Ping from ", pg.Conn.LocalAddr())
		fmt.Print("")
	}
	// 设置超时时间。
	err := pg.SetDeadline()
	if err != nil {
		logger.Error("SetDeadline error: %v", err)
		return err
	}

	// 循环发送ping请求。
	for i := 0; i < count; i++ {
		// 发送ping请求并接收响应。
		r := sendPingMsg(pg.Conn, pg.Data)

		// 如果响应中包含错误，处理错误情况。
		if r.Error != nil {
			// 如果错误是超时错误，则尝试重新连接远程主机。
			if opt, ok := r.Error.(*net.OpError); ok && opt.Timeout() {
				logger.Error("From %s reply: TimeOut", pg.Addr)
				logger.Info("Set Timeout: ", pg.Timeout)
				if err := pg.Dail(); err != nil {
					logger.Error("Not found remote host")
					return r.Error
				}
				return r.Error
			} else {
				if pg.Print {
					// 如果错误不是超时错误，则打印错误信息。
					//fmt.Printf("From %s reply: %s\n", pg.Addr, r.Error)
				}
				continue

			}
		} else {
			// 如果响应正常，则打印响应的详细信息。
			if pg.Print {
				//fmt.Printf("From %s reply: time=%d ttl=%d\n", pg.Addr, r.Time, r.TTL)
			}
		}

		// 等待一段时间后发送下一个ping请求。
		time.Sleep(1e9)
	}
	return nil
}

// PingCount 发送指定数量的ping请求并收集回复。
// 参数count表示要发送的ping请求的数量。
// 返回值reply是一个包含所有ping请求回复的切片。
func (pg *PingSet) PingCount(count int) (reply []Reply) {
	// 尝试连接远程主机，如果失败则打印错误信息并返回空切片。
	if err := pg.Dail(); err != nil {
		logger.Error("Not found remote host")
		return
	}

	// 设置超时时间，这里设为10秒。
	err := pg.SetDeadline()
	if err != nil {
		return nil
	}

	// 循环发送ping请求，共计count次。
	for i := 0; i < count; i++ {
		// 发送ping请求并获取回复。
		r := sendPingMsg(pg.Conn, pg.Data)
		// 将回复添加到结果切片中。
		reply = append(reply, r)
		// 休眠1秒，为下一次发送ping请求做准备。
		time.Sleep(1e9)
	}

	// 返回收集到的所有回复。
	return
}

// New 执行ping操作。
// 它接受一个地址字符串，一个请求编号，和一串字节数据作为输入。
// 返回一个PingSet指针和一个错误值。
// PingSet包含处理后的数据和最终确定的地址。
func New(addr string, timeout time.Duration) (*PingSet, error) {
	// 将请求编号和数据打包成特定格式的消息。
	wb, err := MarshalMsg(8, Data)
	if err != nil {
		// 如果打包过程中出现错误，返回nil和错误信息。
		return nil, err
	}

	// 通过地址查找服务，以确保地址的有效性和可用性。
	// 如果查找失败，返回nil和错误信息。
	addr, err = Lookup(addr)
	if err != nil {
		return nil, err
	}

	// 返回包含打包后数据和有效地址的PingSet结构体指针，以及nil错误。
	return &PingSet{Data: wb, Addr: addr, Timeout: timeout, Req: 8, Print: true}, nil
}
