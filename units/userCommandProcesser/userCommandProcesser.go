package userCommandProcesser

import (
	"SkyWatch/thirdBody/ipProcesser"
	"SkyWatch/thirdBody/portProcesser"
	"flag"
	"time"
)

func (myRoot *CommandReceiver) CommandRev(raw *RawData) *RawData {
	var Args = flag.Args()
	var NArg = flag.NArg()
	var NFlag = flag.NFlag()

	var inputIP = flag.String("ip", "", "Input Your IP")
	var inputIPListc = flag.String("ip_list", "", "Input Your IP_List Path")
	var inputPort = flag.String("port", "", "Input Your port")
	var NoIcmp = flag.Bool("Pn", false, "Do not use ICMP method")
	var UseTopPorts = flag.Bool("top", false, "Scan the most used ports")
	var Thread = flag.Int("T", 200, "Threads for scan")
	var TimeOut = flag.Int("timeout", 800, "timeout (Millisecond)")
	flag.Parse()

	return &RawData{
		Args:        Args,
		NArg:        NArg,
		NFlag:       NFlag,
		IPList:      *inputIPListc,
		IP:          *inputIP,
		Port:        *inputPort,
		NoIcmp:      *NoIcmp,
		UseTopPorts: *UseTopPorts,
		Thread:      *Thread,
		TimeOut:     *TimeOut,
	}
}

func (myRoot *GetIPList) RevCatcher(data *UserCmdProcesser, raw *RawData) {

	if raw.IP != "" {
		rev, _ := ipProcesser.IpGenerater(raw.IP)
		data.IPList = append(data.IPList, rev...)
	}

	if raw.IPList != "" {
		rev, _ := ipProcesser.ReadIPsFromFile(raw.IPList)
		data.IPList = append(data.IPList, rev...)
	}

}

func (myRoot *GetPortList) RevCatcher(data *UserCmdProcesser, raw *RawData) {

	if raw.Port != "" {
		rev, _ := portProcesser.PortGenerater(raw.Port)
		data.Port = append(data.Port, rev...)
	}

}

func (myRoot *GetThread) RevCatcher(data *UserCmdProcesser, raw *RawData) {

	data.Thread = raw.Thread

}

func (myRoot *GetTimeOut) RevCatcher(data *UserCmdProcesser, raw *RawData) {

	data.TimeOut = time.Duration(raw.TimeOut) * time.Millisecond

}

func (myRoot *IfIcmp) RevCatcher(data *UserCmdProcesser, raw *RawData) {

	data.NoIcmp = raw.NoIcmp

}

func (myRoot *IfTopPorts) RevCatcher(data *UserCmdProcesser, raw *RawData) {

	data.UseTopPorts = raw.UseTopPorts

}
