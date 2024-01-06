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
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
)

var (
	ProcName      string // ProcName represents the service name
	Version       string // Version define software version
	CommitHash    string // CommitHash represents the Git commit hash at built time
	BuildDate     string // BuildDate represents the date when this tool was built
	GoVersion     string // GoVersion represents go version
	defaultPrefix string // default prefix path
	isGotest      bool   // Whether go test command
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
	defaultPrefix = wd

	flag.BoolVar(&opts.help, "h", false, "this help")
	flag.BoolVar(&opts.version, "v", false, "show version and exit")
	flag.StringVar(&opts.prefix, "p", defaultPrefix, "set prefix path")
	flag.Var(&opts.signal, "s", "send signal to a process: stop, kill")
	flag.IntVar(&opts.port, "port", 53720, "the port listened on")

	// Whether go test command
	baseName := filepath.Base(os.Args[0])
	if strings.HasSuffix(baseName, ".test") {
		isGotest = true
		ProcName = baseName
	}

	if !isGotest && !flag.Parsed() {
		flag.Parse()
	}

	if opts.help {
		flag.Usage()
		os.Exit(0)
	}
	if opts.version {
		showVersion()
		os.Exit(0)
	}
	if opts.signal == STOPSIG || opts.signal == KILLSIG {
		exit(LogPath())
		os.Exit(0)
	}
}

func Run() {
	writeProcID(LogPath())

	// Set up channel on which to send signal notifications.
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, os.Interrupt)

	addr := fmt.Sprintf("localhost:%d", opts.port)
	server := &http.Server{
		Addr:    addr,
		Handler: nil,
	}

	log.Printf("[cmd.Run] starting server, listen on %s\n", addr)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("[cmd.Run] listen and serve fail, addr=%s, err=%v\n", addr, err)
		}
	}()

	// Block until interrupt signal is received.
	quitSignal := <-quitChan
	log.Println("[cmd.Run] Get signal:", quitSignal)
	server.Close()
}

func Port() int {
	// flagParse()
	return opts.port
}

func PrefixPath() string {
	// flagParse()
	p, err := filepath.Abs(opts.prefix)
	if err != nil {
		log.Printf("[cmd.PrefixPath] get prefix path fail, prefix=%s, err=%v\n", opts.prefix, err)
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
	fmt.Printf("Version\t\t: v%v\nCommit Hash\t: %v\nBuild Date\t: %v\n\n%v\n",
		Version, CommitHash, BuildDate, GoVersion)
}
