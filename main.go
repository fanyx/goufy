package main

import (
	"net"
)

func main() {
	var gbxConnString string
	conn, err := net.Dial("tcp", gbxConnString)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
}
