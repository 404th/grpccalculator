package main

import (
	"log"
	"net"

	"github.com/404th/pkg/adder"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}

	// adder.RegisterAdderServer(s, srv)

	adder.RegisterCalculateServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	reflection.Register(s)

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
