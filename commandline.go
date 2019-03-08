package commandline

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	h        bool
	v        bool
	p        string
	c        string
	procName string
	logPath  string
	port     int
)

func init() {
	procName = filepath.Base(os.Args[0])
	p = filepath.Dir(os.Args[0])
	logPath = p
	port = 32018

	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.StringVar(&p, "p", p, "set `prefix` path")
	flag.StringVar(&c, "c", "etc/conf/nginx.conf", "set configuration `file`")
	flag.StringVar(&logPath, "logpath", logPath, "set the log `path`")
	flag.IntVar(&port, "port", port, "set the service listening `port`")

	flag.Usage = usage

	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(0)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `
Usage: %s [-hv] [-c file] [-p prefix] [-logpath path] [-port port]

Options:
`, procName)
	flag.PrintDefaults()
}

// IsShowVersion return show version flag
func IsShowVersion() bool {
	return v
}

// PrefixPath return prefix path flag
func PrefixPath() string {
	return p
}

// ConfigFile return configuration file flag
func ConfigFile() string {
	return c
}

// LogPath return the log path
func LogPath() string {
	return logPath
}

// Port return the service listening port
func Port() int {
	return port
}
