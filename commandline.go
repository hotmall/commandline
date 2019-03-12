package commandline

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	curPath  = filepath.Dir(os.Args[0])
	procName = filepath.Base(os.Args[0])
	port     = 32018

	h       = flag.Bool("h", false, "this help")
	v       = flag.Bool("v", false, "show version and exit")
	p       = flag.String("p", curPath, "set `prefix` path")
	c       = flag.String("c", "etc/conf/nginx.conf", "set configuration `file`")
	signal  = flag.String("s", "", "send `signal` to the process: stop")
	logPath = flag.String("logpath", curPath, "set the log `path`")
	pport   = flag.Int("port", 32018, "set the service listening `port`")
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `
Usage: %s [-hv] [-c file] [-p prefix] [-logpath path] [-port port]

Options:
`, procName)
	flag.PrintDefaults()
}

// Usage show help
func Usage() {
	// In order to be compatible with the parameters of go-check,
	// flag parsing is moved from init() to each function and parsed when used.
	if !flag.Parsed() {
		flag.Parse()
	}
	flag.Usage()
	os.Exit(0)
}

// ProcName return the process name
func ProcName() string {
	return procName
}

// IsShowHelp return if show help
func IsShowHelp() bool {
	if !flag.Parsed() {
		flag.Parse()
	}
	return *h
}

// IsShowVersion return show version flag
func IsShowVersion() bool {
	if !flag.Parsed() {
		flag.Parse()
	}
	return *v
}

// PrefixPath return prefix path flag
func PrefixPath() string {
	if !flag.Parsed() {
		flag.Parse()
	}
	return *p
}

// ConfigFile return configuration file flag
func ConfigFile() string {
	if !flag.Parsed() {
		flag.Parse()
	}
	return *c
}

// LogPath return the log path
func LogPath() string {
	if !flag.Parsed() {
		flag.Parse()
	}
	return *logPath
}

// Port return the service listening port
func Port() int {
	if !flag.Parsed() {
		flag.Parse()
	}
	return port
}

// Signal return the signal
func Signal() (s SignalType, err error) {
	if !flag.Parsed() {
		flag.Parse()
	}

	err = s.UnmarshalText([]byte(*signal))
	return
}
