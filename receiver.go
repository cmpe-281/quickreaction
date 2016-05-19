package main

import (
    "net/rpc"
    "os"
    "fmt"
    "net"
)

type List int

func (t *List) IncomingData(buf string, reply *int) error {
	fmt.Println(buf)
	return nil
}

// Helper to setup RPC server
func startRPC(port string) {
	// Setup replica server
	
}

func main() {
	p := os.Args[1]

	s, err := net.ResolveTCPAddr("tcp", p)
	if err != nil {
		panic(err)
	}

	i, err := net.ListenTCP("tcp", s)
	if err != nil {
		panic(err)
	}

	list := new(List)
	rpc.Register(list)
	rpc.Accept(i)
}