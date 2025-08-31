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
	aliveHosts     []string
	aliveHostCount int
}

type icmpScanner struct {
	userCommandProcesser.UserCmdProcesser
	AliveHostCount int
	AliveHost      []string
	Task           chan string
	Result         chan string
	Wg             sync.WaitGroup
}
