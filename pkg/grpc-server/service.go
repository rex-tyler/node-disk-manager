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
	"context"
	"fmt"
	"strings"

	protos "github.com/harshthakur9030/node-disk-manager/pkg/protos/ndm"
	ps "github.com/mitchellh/go-ps"
)

//FindISCSIStatus returns the status of the service
func (n *NodeType) FindISCSIStatus(ctx context.Context, null *protos.Null) (*protos.ISCSIStatus, error) {
	processList, _ := ps.Processes()

	var found bool

	for _, p := range processList {
		if strings.Contains(p.Executable(), "iscsi") {
			fmt.Println(p.Executable())
			n.log.Info("ISCSI is running", p.Pid())
			found = true
		}
	}
	if found {
		return &protos.ISCSIStatus{Status: "Running"}, nil
	}
	return &protos.ISCSIStatus{Status: "Not Running"}, nil

}
