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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcID(t *testing.T) {
	assert := assert.New(t)
	ProcName = "TestService"
	pid1 := writeProcID(LogPath())
	pid2 := readProcID(LogPath())
	assert.Equal(pid1, pid2)
}
