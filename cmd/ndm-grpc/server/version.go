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

package server

import (
	protos "github.com/harshthakur9030/node-disk-manager/cmd/ndm-grpc/protos/ndm"
	"github.com/openebs/node-disk-manager/pkg/version"
	"github.com/sirupsen/logrus"

	"context"
)

// VersionType helps in creation of constructor
type VersionType struct {
	log *logrus.Logger
}

// NewVersion is a constructor
func NewVersion(l *logrus.Logger) *VersionType {
	return &VersionType{l}
}

// FindVersion detects returns version and gitCommit
func (v *VersionType) FindVersion(ctx context.Context, null *protos.Null) (*protos.VersionInfo, error) {

	v.log.Info("Print Version")

	return &protos.VersionInfo{Version: version.GetVersion(), GitCommit: version.GetGitCommit()}, nil

}
