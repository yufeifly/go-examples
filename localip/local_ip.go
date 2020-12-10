package localip

import (
	"errors"
	"net"
)

func GetLocalIP() ([]string, error) {
	var localIPs []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIPs = append(localIPs, ipnet.IP.String())
			}
		}
	}
	if localIPs == nil {
		return nil, errors.New("can not find the client ip address")
	}
	return localIPs, nil
}
