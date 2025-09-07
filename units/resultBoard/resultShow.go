package resultBoard

import (
	"github.com/fatih/color"
)

func (root *ResultBoard) ResBoardInit() {
	root.onlineHost = color.New(color.FgHiGreen)

	root.resultBelonging = color.New(color.FgHiBlue).Add(color.Underline)

	root.avaliableServices = color.New(color.FgCyan)

	root.openPort = color.New(color.FgGreen)

	root.underLine = color.New(color.FgCyan)

	root.ErrorLine = color.New(color.FgHiRed)
}

func (root *ResultBoard) StartLine() {
	root.underLine.Printf("=============================================\n")
}

func (root *ResultBoard) ShowOverView(ipaddr string, ports []int) {
	root.resultBelonging.Printf("[IPAddr] < %v > :\n", ipaddr)
	root.openPort.Printf("[+]OpenPort: %v\n", ports)
	root.underLine.Println("------")
}

func (root *ResultBoard) ShowResultDetail(port int, serviceDetail string) {
	root.openPort.Printf("[+]Port %v is open! --> ", port)
	root.avaliableServices.Printf("ServiceType: (%v)\n", serviceDetail)
	//fmt.Println()
}

func (root *ResultBoard) EndLine(count int) {
	root.resultBelonging.Printf("[*]AliveHosts: %v\n", count)
	root.underLine.Printf("=============================================\n")
}

func (root *ResultBoard) NoPortOpen(ipaddr string) {
	root.ErrorLine.Printf("[IPAddr] < %v > : No Openning Port\n", ipaddr)
}
