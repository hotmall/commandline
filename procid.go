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
	"strings"
	"syscall"
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
func readProcID(logPath string) (int, error) {
	pidFile := pidFile(logPath)
	content, err := os.ReadFile(pidFile)
	if err != nil {
		log.Printf("[commandline.readProcID] read pid file fail, pidFile=%s, err=%v\n", pidFile, err)
		return 0, err
	}

	procID, err := strconv.Atoi(string(content))
	if err != nil {
		log.Printf("[commandline.readProcID] strconv procID fail, strProcID=%s, err=%v\n", string(content), err)
		return 0, err
	}
	return procID, nil
}

// removeProcID delete procid file if it records the current process.
// 仅在 PID 文件记录的就是当前进程时才删除，避免误删其他进程写入的 PID 文件。
func removeProcID(logPath string) {
	pidFile := pidFile(logPath)
	content, err := os.ReadFile(pidFile)
	if err != nil {
		return
	}
	if strings.TrimSpace(string(content)) == strconv.Itoa(os.Getpid()) {
		if err := os.Remove(pidFile); err != nil && !os.IsNotExist(err) {
			log.Printf("[commandline.removeProcID] remove pid file fail, pidFile=%s, err=%v\n", pidFile, err)
		}
	}
}

// exit send a SIGINT(2) signal to the process
func exit(logPath string) {
	pid, err := readProcID(logPath)
	if err != nil {
		log.Fatalf("[commandline.exit] read procId from pid file fail, logPath=%s, err=%v\n", logPath, err)
	}
	// 校验 PID 确实属于本服务，避免因 PID 被复用而误向其他进程发送停止信号
	if !isProcessExists(pid) {
		log.Printf("[commandline.exit] the recorded process(%d) is not this service, ignore stop.\n", pid)
		return
	}
	p, err := os.FindProcess(pid)
	if err != nil {
		log.Fatalf("[commandline.exit] find process fail, pid=%d, err=%v\n", pid, err)
	}

	if err = p.Signal(os.Interrupt); err != nil {
		log.Fatalf("[commandline.exit] send interrupt signal fail, pid=%d, err=%v\n", pid, err)
	}
}

// isProcessExists 判断进程是否存在，并校验该进程确实属于本服务。
// 仅依赖 kill(pid,0) 判断存在是不够的：进程退出后其 PID 可能被其他进程复用，
// 导致误判"实例仍存在"而拒绝启动。因此在存活的基础上追加身份校验：
// 读取 /proc/<pid>/cmdline 确认命令行中包含本服务名(ProcName)。
// 若 PID 已被其他进程复用，或进程已变为僵尸（cmdline 为空），均判定为不存在。
func isProcessExists(procId int) bool {
	p, err := os.FindProcess(procId)
	if err != nil {
		log.Printf("[commandline.isProcessExists] find process fail, procId=%d, err=%v\n", procId, err)
		return false
	}
	if err = p.Signal(syscall.Signal(0)); err != nil {
		log.Printf("[commandline.isProcessExists] the process(%d) was already finished, start it.\n", procId)
		return false
	}
	return isProcIDOfService(procId)
}

// isProcIDOfService 校验指定 PID 的进程是否属于本服务。
// Linux 下读取 /proc/<pid>/cmdline（参数以 \0 分隔），检查是否包含服务名。
// 非 Linux 平台（无 /proc）或读取失败时，退化为仅判断存活。
func isProcIDOfService(procId int) bool {
	// 未指定服务名（如 go test 等场景）时，退化为仅判断存活
	if ProcName == "" {
		return true
	}
	cmdline, err := os.ReadFile(filepath.Join("/proc", strconv.Itoa(procId), "cmdline"))
	if err != nil {
		return true
	}
	args := strings.Split(string(cmdline), "\x00")
	return strings.Contains(strings.Join(args, " "), ProcName)
}
