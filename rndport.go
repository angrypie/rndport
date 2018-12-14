package rndport

import (
	"fmt"
	"net"
)

//GetAddress returns free avaliable port or address if template is provided.
func GetAddress(templates ...string) (address string, err error) {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return
	}

	template := "%d"
	if len(templates) == 1 {
		template = templates[0]
	}

	address = fmt.Sprintf(template, listener.Addr().(*net.TCPAddr).Port)
	err = listener.Close()

	return
}
