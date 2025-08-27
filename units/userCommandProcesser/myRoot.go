package userCommandProcesser

type UserCmdCatcher interface {
	RevCatcher(data *UserCmdProcesser, raw *RawData)
}
type UserCmdProcesser struct {
	IPList      []string
	Port        []int
	NoIcmp      bool
	UseTopPorts bool
	Thread      int
	TimeOut     int
}

type RawData struct {
	IPList      string
	IP          string
	Port        string
	NoIcmp      string
	UseTopPorts string
	Thread      int
	TimeOut     int
}

type CommandReceiver struct{}
type GetIPList struct{}

type GetPortList struct{}

type IfIcmp struct{}

type IfTopPorts struct{}

type GetThread struct{}

type GetTimeOut struct{}
