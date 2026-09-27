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
	"strconv"
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

func TestIsProcIDOfService(t *testing.T) {
	assert := assert.New(t)
	pid := os.Getpid()

	old := ProcName
	defer func() { ProcName = old }()

	// 非 Linux 平台（无 /proc）时退化为仅判断存活
	if _, err := os.Stat("/proc"); err != nil {
		ProcName = "definitely_not_the_current_process_marker"
		assert.True(isProcIDOfService(pid))
		return
	}

	// Linux：当前进程命令行不包含该标记 → 判定为不属于本服务（模拟 PID 被复用）
	ProcName = "definitely_not_the_current_process_marker"
	assert.False(isProcIDOfService(pid))

	// 恢复真实服务名后应命中
	ProcName = old
	assert.True(isProcIDOfService(pid))
}

func TestRemoveProcID(t *testing.T) {
	assert := assert.New(t)
	logPath := LogPath()
	pidFile := pidFile(logPath)

	// 写入当前进程 PID，removeProcID 应删除该文件
	writeProcID(logPath)
	removeProcID(logPath)
	_, err := os.Stat(pidFile)
	assert.True(os.IsNotExist(err))

	// 写入其他 PID，removeProcID 不应删除（保护其他进程的 PID 文件）
	other := 99999
	os.WriteFile(pidFile, []byte(strconv.Itoa(other)), 0644)
	removeProcID(logPath)
	content, err := os.ReadFile(pidFile)
	assert.Nil(err)
	assert.Equal(strconv.Itoa(other), string(content))
	os.Remove(pidFile)
}
