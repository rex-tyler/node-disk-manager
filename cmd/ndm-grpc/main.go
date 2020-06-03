package main

import (
	"net"
	"os"

	protos "github.com/harshthakur9030/node-disk-manager/cmd/ndm-grpc/protos/ndm"
	"github.com/harshthakur9030/node-disk-manager/cmd/ndm-grpc/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	log := logrus.New()
	log.Out = os.Stdout

	gs := grpc.NewServer()
	vs := server.NewVersion(*log)

	reflection.Register(gs)

	protos.RegisterVersionServer(gs, vs)

	l, err := net.Listen("tcp", "0.0.0.0:3333")
	if err != nil {
		log.Error("Unable to listen", err)
		os.Exit(1)
	}

	log.Info("Starting server")
	gs.Serve(l)

}
