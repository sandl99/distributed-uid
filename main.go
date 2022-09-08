package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/sandl99/distrbuted-uid/gapi"
	"github.com/sandl99/distrbuted-uid/snowflake"
	"google.golang.org/grpc"
)

func main() {
	snowflake.InfoLogger.Println("Starting the application...")

	var port int
	var node int64
	flag.IntVar(&port, "port", 8000, "port for application.")
	flag.Int64Var(&node, "node", 1, "node id for snowflake.")
	flag.Parse()

	constant, err := snowflake.InitConstant()
	if err != nil {
		snowflake.ErrorLogger.Fatal(err)
	}

	snowflaker, err := snowflake.NewSnowFlake(node, constant)
	if err != nil {
		snowflake.ErrorLogger.Fatal(err)
	}
	snowflake.InfoLogger.Printf("Starting the server on port %d ...\n", port)
	snowflakerServer := gapi.NewSnowFlakeServer(snowflaker)

	server := grpc.NewServer()
	gapi.RegisterSnowFlakeServiceServer(server, snowflakerServer)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		snowflake.ErrorLogger.Fatal("Can not start the server: ", err)
	}

	err = server.Serve(listener)
	if err != nil {
		snowflake.ErrorLogger.Fatal("Can not start the server: ", err)
	}
	snowflake.InfoLogger.Println("Stopping the application...")
}
