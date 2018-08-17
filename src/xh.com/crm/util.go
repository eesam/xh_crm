package crm

import (
	"errors"
	"log"
	"net"
)

func getLocalIp() (string, error) {
	ifc, err := net.InterfaceByName("eth0")
	if err != nil {
		log.Println("Can't find interface eth0.")
		return "", err
	}
	addrs, err := ifc.Addrs()
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	err = errors.New("Can't find interface eth0's ip4 addr.")
	log.Println(err)
	return "", err
}
