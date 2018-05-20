package main

import (
	"fmt"
	"net"
	"os"

	"github.com/amckinney/lsm-go/gen/idl/lsm-go"
	"github.com/amckinney/lsm-go/handler/lsm"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	server := grpc.NewServer()
	handler, err := lsm.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lsmgopb.RegisterLSMGoServer(server, handler)
	server.Serve(listener)
}
