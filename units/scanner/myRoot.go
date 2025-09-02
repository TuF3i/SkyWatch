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
}

type icmpScanner struct {
	//tasks defination
	userCommandProcesser.UserCmdProcesser
	Task chan struct {
		ipAddr string
	}
	Result chan struct {
		ipAddr string
	}
	Wg sync.WaitGroup

	// results defination
	AliveHostCount int
	AliveHost      []string
}

type tcpScanner struct {
	//tasks defination
	userCommandProcesser.UserCmdProcesser
	Task chan struct {
		ipAddr string
		port   int
	}
	Result chan struct {
		ipAddr string
		port   int
	}
	Wg sync.WaitGroup

	//results defination
	openPort map[string][]int
}
