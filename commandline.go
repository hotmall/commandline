package commandline

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	prefix   = curPath()
	procName = filepath.Base(os.Args[0])

	h       = flag.Bool("h", false, "this help")
	v       = flag.Bool("v", false, "show version and exit")
	p       = flag.String("p", prefix, "set `prefix` path")
	c       = flag.String("c", "etc/conf/hot.json", "set configuration `file`")
	signal  = flag.String("s", "", "send `signal` to the process: stop, kill")
	logPath = flag.String("logpath", "var/log", "set the log `path`")
	port    = flag.Int("port", 32018, "set the service listening `port`")
)

func init() {
	flag.Usage = usage
	flag.Parse()

	if *h {
		usage()
		os.Exit(0)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `
Usage: %s [-hv] [-c file] [-p prefix] [-logpath path] [-port port] [-signal stop|kill]

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
	return *h
}

// IsShowVersion return show version flag
func IsShowVersion() bool {
	return *v
}

// PrefixPath return prefix path flag
func PrefixPath() string {
	return *p
}

// ConfigFile return configuration file flag
func ConfigFile() string {
	return *c
}

// LogPath return the log path
func LogPath() string {
	return *logPath
}

// Port return the service listening port
func Port() int {
	return *port
}

// Signal return the signal
func Signal() (s SignalType, err error) {
	err = s.UnmarshalText([]byte(*signal))
	return
}
