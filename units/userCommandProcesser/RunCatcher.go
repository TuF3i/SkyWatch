package userCommandProcesser

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

	for _, r := range UserCmdCatcher {
		r.RevCatcher(&data, raw)
	}

	//fmt.Printf("%v", data)
	return &data
}
