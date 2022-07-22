package shared

import (
	"log"
	"net"
	"strconv"
	"strings"
)

func GetLocalIpAddress(port int) string {
	ipSlice := []string{"127.0.0.1", ":", strconv.Itoa(port)}

	if netInterfaces, err := net.Interfaces(); err != nil {
		log.Println("net.Interfaces failed, err:", err.Error())
	} else {
		for i := 0; i < len(netInterfaces); i++ {
			if (netInterfaces[i].Flags & net.FlagUp) != 0 {
				addrs, _ := netInterfaces[i].Addrs()
				for _, address := range addrs {
					if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
						if ipnet.IP.To4() != nil {
							ipSlice[0] = ipnet.IP.String()
							return strings.Join(ipSlice, "")
						}
					}
				}
			}
		}
	}

	return strings.Join(ipSlice, "")
}
