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
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignal(t *testing.T) {
	assert := assert.New(t)
	items := []struct {
		value string
		want  Signal
		err   error
	}{
		{"stop", STOPSIG, nil},
		{"STOP", STOPSIG, nil},
		{"Stop", STOPSIG, nil},
		{"kill", KILLSIG, nil},
		{"KILL", KILLSIG, nil},
		{"KiLL", KILLSIG, nil},
		{"xxx", 0, errors.New("unkonw signal: xxx")},
	}
	for _, item := range items {
		var s Signal
		err := s.Set(item.value)
		assert.Equal(item.err, err)
		assert.Equal(item.want, s)
	}
}
