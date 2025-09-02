package scanner

import (
	"SkyWatch/thirdBody/icmpScanLib"
	"SkyWatch/units/userCommandProcesser"
	"sync"
)

func (root *icmpScanner) prepareTaskData(data *userCommandProcesser.UserCmdProcesser) {

	/* 初始化ICMP扫描的所有参数 */

	root.IPList = data.IPList
	root.Thread = data.Thread
	root.TimeOut = data.TimeOut
	root.Task = make(chan struct{ ipAddr string })
	root.Result = make(chan struct{ ipAddr string })
	root.AliveHostCount = 0
	root.AliveHost = make([]string, 0)
	root.Wg = sync.WaitGroup{}
}

func (root *icmpScanner) initWorkingThread() {
	for i := 0; i < root.Thread; i++ {
		root.Wg.Add(1)
		go root.worker()
	}
}

func (root *icmpScanner) worker() {
	defer root.Wg.Done()

	for task := range root.Task {
		ifOnline, _ := icmpScanLib.IsHostAlive(task.ipAddr)
		if ifOnline {
			root.Result <- struct{ ipAddr string }{ipAddr: task.ipAddr}
		}
	}
}

func (root *icmpScanner) publishTask() {
	go func() {

		for _, ipAddr := range root.IPList {
			root.Task <- struct{ ipAddr string }{ipAddr: ipAddr}
		}
		close(root.Task)
		//root.Wg.Wait()
		//close(root.Task)
	}()

}

func (root *icmpScanner) waitAllTaskFinish() {
	go func() {
		root.Wg.Wait()
		close(root.Result)
	}()
}

func (root *icmpScanner) Scanner(data *userCommandProcesser.UserCmdProcesser, res *ScannerRoot) {
	root.prepareTaskData(data)
	root.initWorkingThread()
	root.publishTask()
	root.waitAllTaskFinish()

	for result := range root.Result {
		root.AliveHostCount++
		root.AliveHost = append(root.AliveHost, result.ipAddr)
	}

	res.aliveHosts = root.AliveHost
	res.aliveHostCount = root.AliveHostCount

}
