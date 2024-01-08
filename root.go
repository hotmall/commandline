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
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	defaultPort = 53720
)

var (
	ProcName   string // ProcName represents the service name
	Version    string // Version define software version
	CommitHash string // CommitHash represents the Git commit hash at built time
	BuildDate  string // BuildDate represents the date when this tool was built
	GoVersion  string // GoVersion represents go version
	isGotest   bool   // Whether go test command
)

var opts struct {
	help    bool   // this help
	version bool   // show version and exit
	prefix  string // prefix path
	signal  Signal // send signal to a process: stop
	port    int    // the port listened on
}

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("[commandline.init] getwd fail, err = %v\n", err)
	}
	opts.prefix = wd
	opts.port = defaultPort
	prefixUsage := fmt.Sprintf("set prefix path (default: %s)", opts.prefix)
	portUsage := fmt.Sprintf("set the port listened on (default: %d)", opts.port)

	flag.BoolFunc("h", "show this help", func(help string) error {
		if opts.help, err = strconv.ParseBool(help); err != nil {
			return err
		}
		if opts.help {
			flag.Usage()
			os.Exit(0)
		}
		return nil
	})
	flag.BoolFunc("v", "show version and exit", func(version string) error {
		if opts.version, err = strconv.ParseBool(version); err != nil {
			return err
		}
		if opts.version {
			showVersion()
			os.Exit(0)
		}
		return nil
	})
	flag.Func("p", prefixUsage, func(prefix string) error {
		opts.prefix = prefix
		return nil
	})
	flag.Func("s", "send signal to a process: stop", func(signal string) error {
		if err = opts.signal.Set(signal); err != nil {
			return (err)
		}
		if opts.signal == STOPSIG || opts.signal == KILLSIG {
			exit(LogPath())
			os.Exit(0)
		}
		return nil
	})
	flag.Func("port", portUsage, func(port string) error {
		v, err := strconv.ParseInt(port, 0, strconv.IntSize)
		if err != nil {
			return err
		}
		opts.port = int(v)
		return nil
	})

	// Whether go test command
	baseName := filepath.Base(os.Args[0])
	if strings.HasSuffix(baseName, ".test") || strings.HasPrefix(baseName, "__debug_bin") {
		isGotest = true
		ProcName = baseName
	}

	if !isGotest && !flag.Parsed() {
		flag.Parse()
	}
}

func Run() {
	// 判断进程是否已经启动，如果已经启动，则退出
	logPath := LogPath()
	if pid, err := readProcID(logPath); err == nil {
		if isProcessExists(pid) {
			log.Fatalf("[commandline.Run] the process(%d) is exists, exit.\n", pid)
		}
	}
	// 进程实例不存在，再启动
	writeProcID(logPath)

	// Set up channel on which to send signal notifications.
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, os.Interrupt)

	addr := fmt.Sprintf("localhost:%d", opts.port)
	server := &http.Server{
		Addr:    addr,
		Handler: nil,
	}

	log.Printf("[commandline.Run] starting server, listen on %s\n", addr)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("[commandline.Run] listen and serve fail, addr=%s, err=%v\n", addr, err)
		}
	}()

	// Block until interrupt signal is received.
	quitSignal := <-quitChan
	log.Println("[commandline.Run] Get signal:", quitSignal)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("[commandline.Run] shutdown server fail, err=%v\n", err)
	}
}

func Port() int {
	return opts.port
}

func PrefixPath() string {
	p, err := filepath.Abs(opts.prefix)
	if err != nil {
		log.Printf("[commandline.PrefixPath] get prefix path fail, prefix=%s, err=%v\n", opts.prefix, err)
		p = filepath.Clean(opts.prefix)
	}
	return p
}

func SetPrefixPath(prefix string) {
	opts.prefix = prefix
}

func LogPath() string {
	return filepath.Join(PrefixPath(), "var", "log")
}

func ConfigFile() string {
	return filepath.Join(PrefixPath(), "etc", "conf", "Hotfile")
}

func IsGotest() bool {
	return isGotest
}

func showVersion() {
	var commit string
	if len(CommitHash) > 0 {
		commit = fmt.Sprintf("(commit: %s)", CommitHash)
	}
	const govprefix = "go version "
	var gov string
	if strings.HasPrefix(GoVersion, govprefix) {
		gov = strings.TrimPrefix(GoVersion, govprefix)
	}
	fmt.Printf("%s version %s%s builded by %s at %s\n", ProcName, Version, commit, gov, BuildDate)
}
