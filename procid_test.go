// Copyright © 2024 The Hot Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commandline

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcID(t *testing.T) {
	assert := assert.New(t)
	pid1 := writeProcID(LogPath())
	pid2, err := readProcID(LogPath())
	assert.Nil(err)
	assert.Equal(pid1, pid2)
}

func TestIsProcessExists(t *testing.T) {
	assert := assert.New(t)
	items := []struct {
		pid  int
		want bool
	}{
		{100, false},
		{200, false},
		{300, false},
		{400, false},
	}

	for _, item := range items {
		ret := isProcessExists(item.pid)
		assert.Equal(item.want, ret)
	}

	pid := os.Getpid()
	ret := isProcessExists(pid)
	assert.True(ret)
}
