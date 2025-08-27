package portProcesser

import (
	"fmt"
	"strconv"
	"strings"
)

func parsePortSequence(portSeq string) ([]int, error) {
	parts := strings.Split(portSeq, "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("无效的端口范围格式")
	}

	start, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return nil, fmt.Errorf("无效的起始端口")
	}

	end, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return nil, fmt.Errorf("无效的结束端口")
	}

	if start > end {
		return nil, fmt.Errorf("起始端口不能大于结束端口")
	}

	if start < 1 || end > 65535 {
		return nil, fmt.Errorf("端口号超出范围 (1-65535)")
	}

	var ports []int
	for i := start; i <= end; i++ {
		ports = append(ports, i)
	}

	return ports, nil
}

func PortGenerater(portRange string) ([]int, error) {
	var ports []int

	// 处理逗号分隔的端口
	if strings.Contains(portRange, ",") {
		portStrs := strings.Split(portRange, ",")
		for _, portStr := range portStrs {
			if strings.Contains(portStr, "-") {
				rangePorts, err := parsePortSequence(portStr)
				if err != nil {
					return nil, err
				}
				ports = append(ports, rangePorts...)
			} else {
				port, err := strconv.Atoi(strings.TrimSpace(portStr))
				if err != nil {
					return nil, fmt.Errorf("无效的端口: %s", portStr)
				}
				if port < 1 || port > 65535 {
					return nil, fmt.Errorf("端口号超出范围 (1-65535): %d", port)
				}
				ports = append(ports, port)
			}
		}
		return ports, nil
	}

	// 处理端口范围
	if strings.Contains(portRange, "-") {
		return parsePortSequence(portRange)
	}

	// 单个端口
	port, err := strconv.Atoi(portRange)
	if err != nil {
		return nil, fmt.Errorf("无效的端口: %s", portRange)
	}
	if port < 1 || port > 65535 {
		return nil, fmt.Errorf("端口号超出范围 (1-65535): %d", port)
	}
	return []int{port}, nil
}
