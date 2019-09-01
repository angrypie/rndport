package rndport

import (
	"fmt"
	"net"
)

//GetAddress returns template string with free available port. Template is optional (default "%d").
func GetAddress(templates ...string) (address string, err error) {
	port, err := GetPort()
	if err != nil {
		return
	}

	template := "%d"
	if len(templates) == 1 {
		template = templates[0]
	}

	address = fmt.Sprintf(template, port)

	return
}

//GetPort returns free available port.
func GetPort() (port int, err error) {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return
	}
	port = listener.Addr().(*net.TCPAddr).Port
	err = listener.Close()

	return

}
