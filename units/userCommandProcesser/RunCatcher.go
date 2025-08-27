package userCommandProcesser

import "fmt"

func RunCatcher() {
	UserCmdCatcher := []UserCmdCatcher{
		&GetIPList{},
		&GetPortList{},
	}

	data := UserCmdProcesser{}
	cmdRev := CommandReceiver{}

	raw := cmdRev.CommandRev(&RawData{})
	for _, r := range UserCmdCatcher {
		r.RevCatcher(&data, raw)
	}

	fmt.Printf("%v", data)
}
