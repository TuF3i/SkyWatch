package userCommandProcesser

import (
	"SkyWatch/thirdBody/ipProcesser"
	"SkyWatch/thirdBody/portProcesser"
	"flag"
	"os"
	"time"
)

func (myRoot *CommandReceiver) CommandRev(raw *RawData) *RawData {
	var Args = flag.Args()
	var NArg = flag.NArg()
	var NFlag = flag.NFlag()
	var IfArgs = os.Args

	var inputIP = flag.String("ip", "", "Input Your IP")
	var inputIPListc = flag.String("ip_list", "", "Input Your IP_List Path")
	var inputPort = flag.String("port", "", "Input Your port")
	var NoIcmp = flag.Bool("Pn", false, "Do not use ICMP method")
	var UseTopPorts = flag.Bool("top", false, "Scan the most used ports")
	var Thread = flag.Int("T", 200, "Threads for scan")
	var TimeOut = flag.Int("timeout", 800, "timeout (Millisecond)")
	flag.Parse()

	return &RawData{
		IfArgs:      IfArgs,
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

	GetCommonPorts := func() []int {
		return []int{
			20, 21, 22, 23, 25, 53, 67, 68, 69, 80,
			110, 119, 123, 135, 137, 138, 139, 143, 161, 162,
			389, 443, 445, 465, 514, 515, 587, 631, 636, 993,
			995, 1080, 1194, 1433, 1434, 1521, 1723, 1863, 2049, 2082,
			2083, 2086, 2087, 2095, 2096, 2222, 2375, 2376, 3000, 3128,
			3306, 3389, 3690, 4000, 4040, 4430, 4500, 4567, 4662, 4672,
			4899, 5000, 5001, 5004, 5005, 5050, 5060, 5070, 5100, 5190,
			5222, 5223, 5269, 5353, 5355, 5432, 5500, 5631, 5632, 5666,
			5800, 5900, 6000, 6001, 6379, 6566, 6665, 6666, 6667, 6668,
			6669, 6679, 6697, 6881, 6882, 6883, 6884, 6885, 6886, 6887,
			6888, 6889, 6890, 6891, 6901, 6969, 6970, 7212, 7648, 8000,
		}
	}

	if data.UseTopPorts {
		data.UseTopPorts = raw.UseTopPorts
		data.Port = append(data.Port, GetCommonPorts()...)
	}
}
