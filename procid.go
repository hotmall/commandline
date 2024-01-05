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
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// pidFile return procid file
func pidFile(logPath string) (pidFile string) {
	if err := os.MkdirAll(logPath, os.ModePerm); err != nil {
		log.Fatalf("[commandline.pidFile] make all dir fail, logPath=%s, err=%v\n", logPath, err)
	}
	pidFile = filepath.Join(logPath, ProcName+".pid")
	return
}

// writeProcID write procID into xx.pid file
func writeProcID(logPath string) int {
	pidFile := pidFile(logPath)
	pid := os.Getpid()
	str := strconv.Itoa(pid)
	err := os.WriteFile(pidFile, []byte(str), 0644)
	if err != nil {
		log.Fatalf("[commandline.writeProcID] write pid file fail, pidFile=%s, err=%v\n", pidFile, err)
	}
	return pid
}

// readProcID read procID from xx.pid file
func readProcID(logPath string) int {
	pidFile := pidFile(logPath)

	content, err := os.ReadFile(pidFile)
	if err != nil {
		log.Fatalf("[commandline.readProcID] read pid file fail, pidFile=%s, err=%v\n", pidFile, err)
	}

	procID, err := strconv.Atoi(string(content))
	if err != nil {
		log.Fatalf("[commandline.readProcID] strconv procID fail, strProcID=%s, err=%v\n", string(content), err)
	}
	return procID
}

// exit send a SIGINT(2) signal to the process
func exit(logPath string) {
	pid := readProcID(logPath)
	p, err := os.FindProcess(pid)
	if err != nil {
		log.Fatalf("[commandline.exit] find process fail, pid=%d, err=%v\n", pid, err)
	}

	if err = p.Signal(os.Interrupt); err != nil {
		log.Fatalf("[commandline.exit] send interrupt signal fail, pid=%d, err=%v\n", pid, err)
	}
}
