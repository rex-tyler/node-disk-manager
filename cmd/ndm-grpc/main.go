/*
Copyright 2019 The OpenEBS Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
	vs := server.NewVersion(log)
	ns := server.NewNode(log)

	reflection.Register(gs)

	protos.RegisterVersionServer(gs, vs)

	protos.RegisterServiceInfoServer(gs, ns)

	l, err := net.Listen("tcp", "0.0.0.0:3333")
	if err != nil {
		log.Error("Unable to listen", err)
		os.Exit(1)
	}

	log.Info("Starting server")
	gs.Serve(l)

}
