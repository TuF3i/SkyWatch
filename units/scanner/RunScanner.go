package scanner

import (
	"SkyWatch/units/userCommandProcesser"
)

func RunScanner(data *userCommandProcesser.UserCmdProcesser) *ScannerRoot {
	root := ScannerRoot{}

	Scanner := []Scanner{
		&icmpScanner{},
		&tcpScanner{},
		&serviceScanner{},
	}

	for _, r := range Scanner {
		r.Scanner(data, &root)
	}

	return &root

}
