package scanner

import (
	"SkyWatch/thirdBody/tcpScanLib"
	"SkyWatch/units/userCommandProcesser"
	"sync"
)

func (root *tcpScanner) prepareTaskData(data *userCommandProcesser.UserCmdProcesser, res *ScannerRoot) {

	/* 初始化ICMP扫描的所有参数 */

	root.IPList = res.aliveHosts
	root.Port = data.Port
	root.Thread = data.Thread
	root.TimeOut = data.TimeOut
	root.Task = make(chan tcpTaskUnity)
	root.Result = make(chan tcpResultUnity)
	root.openPort = make(map[string][]int)
	root.Wg = sync.WaitGroup{}
}

func (root *tcpScanner) initWorkingThread() {
	for i := 0; i < root.Thread; i++ {
		root.Wg.Add(1)
		go root.worker()
	}
}

func (root *tcpScanner) worker() {
	defer root.Wg.Done()

	for task := range root.Task {
		ifOpen, _ := tcpScanLib.TCPPortScan(task.ipAddr, task.port, root.TimeOut)
		if ifOpen {
			root.Result <- tcpResultUnity{ipAddr: task.ipAddr, port: task.port}
		}
	}
}

func (root *tcpScanner) publishTask() {
	go func() {
		for _, ipAddr := range root.IPList {
			for _, port := range root.Port {
				root.Task <- tcpTaskUnity{ipAddr: ipAddr, port: port}
			}
		}
		close(root.Task)
	}()

}

func (root *tcpScanner) waitAllTaskFinish() {
	go func() {
		root.Wg.Wait()
		close(root.Result)
	}()
}

func (root *tcpScanner) Scanner(data *userCommandProcesser.UserCmdProcesser, res *ScannerRoot) {
	root.prepareTaskData(data, res)
	root.initWorkingThread()
	root.publishTask()
	root.waitAllTaskFinish()

	for result := range root.Result {
		root.openPort[result.ipAddr] = append(root.openPort[result.ipAddr], result.port)
	}

	res.openPort = root.openPort

}
