package userCommandProcesser

import "time"

type UserCmdCatcher interface {
	RevCatcher(data *UserCmdProcesser, raw *RawData)
}
type UserCmdProcesser struct {
	IPList      []string
	Port        []int
	NoIcmp      bool
	UseTopPorts bool
	Thread      int
	TimeOut     time.Duration
}

type RawData struct {
	Args        []string
	NArg        int
	NFlag       int
	IPList      string
	IP          string
	Port        string
	NoIcmp      bool
	UseTopPorts bool
	Thread      int
	TimeOut     int
	IfArgs      []string
}

type CommandReceiver struct{}
type GetIPList struct{}

type GetPortList struct{}

type IfIcmp struct{}

type IfTopPorts struct{}

type GetThread struct{}

type GetTimeOut struct{}
