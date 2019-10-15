package commandline

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	prefix         = curPath()
	procName       = filepath.Base(os.Args[0])
	defaultConfig  = strings.Join([]string{prefix, "etc", "conf", "hot.json"}, string(filepath.Separator))
	defaultLogPath = strings.Join([]string{prefix, "var", "log"}, string(filepath.Separator))

	h       = flag.Bool("h", false, "this help")
	v       = flag.Bool("v", false, "show version and exit")
	p       = flag.String("p", prefix, "set `prefix` path")
	c       = flag.String("c", defaultConfig, "set configuration `file`")
	signal  = flag.String("s", "", "send `signal` to the process: stop, kill")
	logPath = flag.String("logpath", defaultLogPath, "set the log `path`")
	port    = flag.Int("port", 32018, "set the service listening `port`")
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `
Usage: %s [-hv] [-c file] [-p prefix] [-s stop|kill] [-logpath path] [-port port]

Options:
`, procName)
	flag.PrintDefaults()
}

// Usage show help
func Usage() {
	flag.Usage()
}

func curPath() string {
	d := filepath.Dir(os.Args[0])
	d += string(filepath.Separator) + ".."
	d = filepath.Clean(d)
	return d
}

// ProcName return the process name
func ProcName() string {
	return procName
}

// CurPath return the current path
func CurPath() string {
	return prefix
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
	if *p != prefix && *c == defaultConfig {
		*c = strings.Join([]string{*p, "etc", "conf", "hot.json"}, string(filepath.Separator))
	}
	return *c
}

// LogPath return the log path
func LogPath() string {
	if !flag.Parsed() {
		flag.Parse()
	}
	if *p != prefix && *logPath == defaultLogPath {
		*logPath = strings.Join([]string{*p, "var", "log"}, string(filepath.Separator))
	}
	return *logPath
}

// Port return the service listening port
func Port() int {
	if !flag.Parsed() {
		flag.Parse()
	}
	return *port
}

// Signal return the signal
func Signal() (s SignalType, err error) {
	if !flag.Parsed() {
		flag.Parse()
	}
	err = s.UnmarshalText(*signal)
	return
}
