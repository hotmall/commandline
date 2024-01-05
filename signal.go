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
	"fmt"
	"strings"
)

// A user-defined flag type
type Signal int8

const (
	STOPSIG Signal = 0x2
	KILLSIG Signal = 0x9
)

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (i *Signal) String() string {
	var v string
	switch *i {
	case STOPSIG:
		v = "stop"
	case KILLSIG:
		v = "kill"
	default:
		v = fmt.Sprintf("unknow signal: %d", *i)
	}
	return v
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
func (i *Signal) Set(value string) error {
	var err error
	s := strings.ToLower(value)
	switch s {
	case "stop":
		*i = STOPSIG
	case "kill":
		*i = KILLSIG
	default:
		err = fmt.Errorf("unkonw signal: %s", value)
	}
	return err
}
