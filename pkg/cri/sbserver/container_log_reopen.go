/*
   Copyright The containerd Authors.

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

package sbserver

import (
	"context"
	"errors"
	"fmt"

	runtime "k8s.io/cri-api/pkg/apis/runtime/v1"
)

// ReopenContainerLog asks the cri plugin to reopen the stdout/stderr log file for the container.
// This is often called after the log file has been rotated.
func (c *criService) ReopenContainerLog(ctx context.Context, r *runtime.ReopenContainerLogRequest) (*runtime.ReopenContainerLogResponse, error) {
	container, err := c.containerStore.Get(r.GetContainerId())
	if err != nil {
		return nil, fmt.Errorf("an error occurred when try to find container %q: %w", r.GetContainerId(), err)
	}

	if container.Status.Get().State() != runtime.ContainerState_CONTAINER_RUNNING {
		return nil, errors.New("container is not running")
	}

	// Create new container logger and replace the existing ones.
	stdoutWC, stderrWC, err := c.createContainerLoggers(container.LogPath, container.Config.GetTty())
	if err != nil {
		return nil, err
	}
	oldStdoutWC, oldStderrWC := container.IO.AddOutput("log", stdoutWC, stderrWC)
	if oldStdoutWC != nil {
		oldStdoutWC.Close()
	}
	if oldStderrWC != nil {
		oldStderrWC.Close()
	}
	return &runtime.ReopenContainerLogResponse{}, nil
}
