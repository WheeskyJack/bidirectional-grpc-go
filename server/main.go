package main

import (
	"bidirectional-grpc-go/max"
	pb "bidirectional-grpc-go/proto"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var l4g = setLog(os.Stdout, "server")

func setLog(out io.Writer, prefix string) *log.Logger {
	l := log.New(out, prefix, log.Lshortfile)
	l.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	return l
}

func main() {
	startServer("127.0.0.1", "9090")
}

func startServer(addr, port string) {
	defer l4g.Println("exiting...")
	serveraddr := addr + ":" + port
	l4g.Println("Starting GRPC server on " + serveraddr)
	lis, err := net.Listen("tcp4", serveraddr)
	if err != nil {
		l4g.Printf("unable to listen on %v : %v", serveraddr, err)
		return
	}
	s := grpc.NewServer()
	m := new(max.Serve)
	pb.RegisterFindServer(s, m)
	// Register registers the server reflection service on the given gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		l4g.Printf("Failed to serve: %v", err)
	}
}
