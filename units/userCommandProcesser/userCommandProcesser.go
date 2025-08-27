package userCommandProcesser

import (
	"SkyWatch/thirdBody/ipProcesser"
	"SkyWatch/thirdBody/portProcesser"
	"flag"
)

func (myRoot *CommandReceiver) CommandRev(raw *RawData) *RawData {
	var inputIP = flag.String("ip", "", "Input Your IP")
	var inputIPListc = flag.String("ip_list", "", "Input Your IP_List Path")
	var inputPort = flag.String("port", "", "Input Your port")
	var NoIcmp = flag.String("Pn", "", "Do not use ICMP method")
	var UseTopPorts = flag.String("top", "", "Scan the most used ports")
	var Thread = flag.Int("T", 200, "Threads for scan")
	var TimeOut = flag.Int("timeout", 1, "timeout (default: 1 second)")
	flag.Parse()

	return &RawData{
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
