package main

import (
	"SkyWatch/units/resultBoard"
	"SkyWatch/units/scanner"
	"SkyWatch/units/userCommandProcesser"
)

func main() {
	banner := resultBoard.ResultBoard{}
	banner.Banner()

	source := userCommandProcesser.RunCatcher()
	res := scanner.RunScanner(source)
	resultBoard.ShowRes(res)
	//fmt.Printf("%v", source)
	//fmt.Printf("%v", res)
}
