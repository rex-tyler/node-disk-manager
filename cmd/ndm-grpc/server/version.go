package server

import (
	protos "github.com/harshthakur9030/node-disk-manager/cmd/ndm-grpc/protos/ndm"
	"github.com/openebs/node-disk-manager/pkg/version"
	"github.com/sirupsen/logrus"

	"context"
)

// VersionType helps in creation of constructor
type VersionType struct {
	log logrus.Logger
}

// NewVersion is a constructor
func NewVersion(l logrus.Logger) *VersionType {
	return &VersionType{l}
}

// FindVersion detects returns version and gitCommit
func (v *VersionType) FindVersion(ctx context.Context, null *protos.Null) (*protos.VersionInfo, error) {

	v.log.Info("Print Version")

	return &protos.VersionInfo{Version: version.GetVersion(), GitCommit: version.GetGitCommit()}, nil

}
