package scanner

import (
	"SkyWatch/units/userCommandProcesser"
)

func RunScanner(data *userCommandProcesser.UserCmdProcesser) *ScannerRoot {
	root := ScannerRoot{}
	var Scanner_ []Scanner
	normalFunc := []Scanner{
		&tcpScanner{},
		&serviceScanner{},
	}

	if data.NoIcmp {
		Scanner_ = append(Scanner_, normalFunc...)
		root.AliveHosts = data.IPList
	} else {
		Scanner_ = append([]Scanner{&icmpScanner{}}, normalFunc...)
	}

	for _, r := range Scanner_ {
		r.Scanner(data, &root)
	}

	return &root

}
