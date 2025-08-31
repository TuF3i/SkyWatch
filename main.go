package main

import (
	"SkyWatch/units/scanner"
	"SkyWatch/units/userCommandProcesser"
	"fmt"
)

func main() {
	source := userCommandProcesser.RunCatcher()
	res := scanner.RunScanner(source)
	fmt.Printf("%v", source)
	fmt.Printf("%v", res)
}
