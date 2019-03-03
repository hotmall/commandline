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
)

func init() {
	procName = filepath.Base(os.Args[0])
	p = filepath.Dir(os.Args[0])

	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.StringVar(&p, "p", p, "set `prefix` path")
	flag.StringVar(&c, "c", "etc/conf/nginx.conf", "set configuration `file`")

	flag.Usage = usage

	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(0)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `
Usage: %s [-hv] [-c file] [-p prefix]

Options:
`, procName)
	flag.PrintDefaults()
}

// IsShowVersion return show version flag
func IsShowVersion() bool {
	return v
}

// GetPrefixPath return prefix path flag
func GetPrefixPath() string {
	return p
}

// GetConfigFile return configuration file flag
func GetConfigFile() string {
	return c
}
