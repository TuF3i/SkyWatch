package scanner

import (
	"SkyWatch/thirdBody/serviceScanLib"
	"SkyWatch/units/userCommandProcesser"
	"fmt"
	"sync"
)

func (root *serviceScanner) prepareTaskData(data *userCommandProcesser.UserCmdProcesser, res *ScannerRoot) {

	/* 初始化Service扫描的所有参数 */

	root.openPort = res.openPort
	root.Thread = data.Thread
	root.TimeOut = data.TimeOut
	root.Task = make(chan serviceTaskUnity)
	root.Result = make(chan serviceResultUnity)
	root.Wg = sync.WaitGroup{}
	root.serviceDetails = make(map[string][]serviceMid)
}

func (root *serviceScanner) initWorkingThread() {
	for i := 0; i < root.Thread; i++ {
		root.Wg.Add(1)
		go root.worker()
	}
}

func (root *serviceScanner) worker() {
	defer root.Wg.Done()

	for task := range root.Task {
		serviceInfo := vscan.GetProbes(fmt.Sprintf("%v:%v", task.ipAddr, task.port))
		root.Result <- serviceResultUnity{ipAddr: task.ipAddr, port: task.port, serviceInfo: serviceInfo}
	}
}

func (root *serviceScanner) publishTask() {
	go func() {
		for ipAddr, ports := range root.openPort {
			for _, port := range ports {
				root.Task <- serviceTaskUnity{ipAddr: ipAddr, port: port}
			}
		}
		close(root.Task)
	}()

}

func (root *serviceScanner) waitAllTaskFinish() {
	go func() {
		root.Wg.Wait()
		close(root.Result)
	}()
}

func (root *serviceScanner) Scanner(data *userCommandProcesser.UserCmdProcesser, res *ScannerRoot) {
	root.prepareTaskData(data, res)
	root.initWorkingThread()
	root.publishTask()
	root.waitAllTaskFinish()

	for result := range root.Result {
		root.serviceDetails[result.ipAddr] = append(root.serviceDetails[result.ipAddr], serviceMid{port: result.port, serviceInfo: result.serviceInfo})
	}

	res.serviceDetails = root.serviceDetails

}
