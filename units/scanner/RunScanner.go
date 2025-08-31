package scanner

import (
	"SkyWatch/units/userCommandProcesser"
)

func RunScanner(data *userCommandProcesser.UserCmdProcesser) *ScannerRoot {
	Scanner := []Scanner{
		&icmpScanner{},
	}

	root := ScannerRoot{}

	for _, r := range Scanner {
		r.Scanner(data, &root)
	}

	return &root

}
