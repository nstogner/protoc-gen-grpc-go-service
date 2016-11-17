package main

import (
	"log"
	"net"
	"os"

	"github.com/nstogner/protoc-gen-grpc-goservice/example"

	"google.golang.org/grpc"
)

func main() {
	port := "4001"
	if p, ok := os.LookupEnv("PORT"); ok {
		port = p
	}

	log.Fatal(listen(port))
}

func listen(port string) error {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	example.RegisterExampleServiceServer(s, &service{})

	return s.Serve(l)
}
