package scanner

import (
	"SkyWatch/units/userCommandProcesser"
	"sync"
)

//type UserCmdProcesser userCommandProcesser.UserCmdProcesser

type Scanner interface {
	Scanner(data *userCommandProcesser.UserCmdProcesser, res *ScannerRoot)
}

type ScannerRoot struct {
	// icmpScanner Results
	aliveHosts     []string
	aliveHostCount int

	// tcpScanner Results
	openPort map[string][]int

	// serviceScanner Result
	serviceDetails map[string][]serviceMid
}

type icmpTaskUnity struct {
	ipAddr string
}
type icmpResultUnity struct {
	ipAddr string
}

type tcpTaskUnity struct {
	ipAddr string
	port   int
}
type tcpResultUnity struct {
	ipAddr string
	port   int
}

type serviceTaskUnity struct {
	ipAddr string
	port   int
}
type serviceResultUnity struct {
	ipAddr      string
	serviceInfo string
	port        int
}
type serviceMid struct {
	port        int
	serviceInfo string
}

type icmpScanner struct {
	//tasks defination
	userCommandProcesser.UserCmdProcesser
	Task   chan icmpTaskUnity
	Result chan icmpResultUnity
	Wg     sync.WaitGroup

	// results defination
	AliveHostCount int
	AliveHost      []string
}

type tcpScanner struct {
	//tasks defination
	userCommandProcesser.UserCmdProcesser
	Task   chan tcpTaskUnity
	Result chan tcpResultUnity
	Wg     sync.WaitGroup

	//results defination
	openPort map[string][]int
}

type serviceScanner struct {
	//task defination
	userCommandProcesser.UserCmdProcesser
	Task     chan serviceTaskUnity
	Result   chan serviceResultUnity
	Wg       sync.WaitGroup
	openPort map[string][]int

	//results defination
	serviceDetails map[string][]serviceMid
}
