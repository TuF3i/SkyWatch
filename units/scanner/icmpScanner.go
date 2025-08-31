package scanner

import (
	"SkyWatch/thirdBody/icmpScanLib"
	"SkyWatch/units/userCommandProcesser"
	"sync"
)

func (root *icmpScanner) prepareTaskData(data *userCommandProcesser.UserCmdProcesser) {
	root.IPList = data.IPList
	root.Thread = data.Thread
	root.TimeOut = data.TimeOut
	root.Task = make(chan string, len(root.IPList))
	root.Result = make(chan string, len(root.IPList))
	root.AliveHostCount = 0
	root.AliveHost = make([]string, 0)
	root.Wg = sync.WaitGroup{}
}

func (root *icmpScanner) initWorkingThread() {
	go func() {
		for i := 0; i < root.Thread; i++ {
			root.Wg.Add(1)
			go root.worker()
			//go worker(tasks, results, &wg, *timeout)
		}
	}()
}

func (root *icmpScanner) worker() {
	defer root.Wg.Done()

	for ipaddr := range root.Task {
		ifOnline, _ := icmpScanLib.IsHostAlive(ipaddr)
		if ifOnline {
			root.Result <- ipaddr
		}
	}
}

func (root *icmpScanner) publishTask() {
	go func() {
		for _, ipAddr := range root.IPList {
			root.Task <- ipAddr
		}
		close(root.Task)
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

	//time.Sleep(time.Second)
	for ipaddr := range root.Result {
		root.AliveHostCount++
		root.AliveHost = append(root.AliveHost, ipaddr)
	}

	res.aliveHosts = root.AliveHost
	res.aliveHostCount = root.AliveHostCount
}
