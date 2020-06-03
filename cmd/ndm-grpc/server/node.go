package server

import (
	"context"
	"os"

	protos "github.com/harshthakur9030/node-disk-manager/cmd/ndm-grpc/protos/ndm"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type NodeType struct {
	log logrus.Logger
}

// NewNode is a constructor
func NewNode(l logrus.Logger) *NodeType {
	return &NodeType{l}
}

// FindNodeName is used to find the name of the worker node NDM is deployed on
func (n *NodeType) FindNodeName(ctx context.Context, in *protos.Null, opts ...grpc.CallOption) (*protos.NodeName, error) {
	nodeName := os.Getenv("NODE_NAME")
	return &protos.NodeName{NodeName: nodeName}, nil

}

//FindServiceStatus returns the status of the service
func (n *NodeType) FindServiceStatus(ctx context.Context, in *protos.ServiceName, opts ...grpc.CallOption) (*protos.ServiceStatus, error) {
	return &protos.ServiceStatus{Status: "Running"}, nil
}
