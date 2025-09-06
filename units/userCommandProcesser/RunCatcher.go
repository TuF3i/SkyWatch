package userCommandProcesser

import "fmt"

func RunCatcher() *UserCmdProcesser {
	UserCmdCatcher := []UserCmdCatcher{
		&GetIPList{},
		&GetPortList{},
		&GetThread{},
		&GetTimeOut{},
		&IfIcmp{},
		&IfTopPorts{},
	}

	data := UserCmdProcesser{}
	cmdRev := CommandReceiver{}
	raw := cmdRev.CommandRev(&RawData{})

	if len(raw.IfArgs) == 0 {
		fmt.Println("Type \"SkyWatch -h\" for help!")
	}

	for _, r := range UserCmdCatcher {
		r.RevCatcher(&data, raw)
	}

	//fmt.Printf("%v", data)
	return &data
}
