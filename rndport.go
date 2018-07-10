package rndport

import (
	"net"
	"strconv"
)

func GetAddress(template ...string) (port string, err error) {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return
	}
	port = strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)
	err = listener.Close()

	return port, err
}
