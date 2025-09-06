package resultBoard

import "SkyWatch/units/scanner"

func ShowRes(data *scanner.ScannerRoot) {
	root := ResultBoard{}
	root.ResBoardInit()
	root.StartLine()
	for _, aliveIp := range data.AliveHosts {
		if len(data.OpenPort[aliveIp]) != 0 {
			root.ShowOverView(aliveIp, data.OpenPort[aliveIp])
			for _, details_ := range data.ServiceDetails[aliveIp] {
				root.ShowResultDetail(details_.Port, details_.ServiceInfo)
			}
		} else {
			root.NoPortOpen(aliveIp)
		}
	}
	root.EndLine(data.AliveHostCount)
}
