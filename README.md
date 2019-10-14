# commandline

## Installation

You can use `go get` to get the latest version:

```shell
go get -u github.com/mallbook/commandline
```

## Quick Start

### Show Help

```shell
Usage: GitlabDriver [-hv] [-c file] [-p prefix] [-s stop|kill] [-logpath path] [-port port]

Options:
  -c file
      set configuration file (default "./etc/conf/hot.json")
  -h  this help
  -logpath path
      set the log path (default "./var/log")
  -p prefix
      set prefix path (default ".")
  -port port
      set the service listening port (default 32018)
  -s signal
      send signal to the process: stop, kill
  -v  show version and exit
```

### Usage

```go
// Print Version
if commandline.IsShowVersion() {
    fmt.Println("version: 1.0.1, build: 2019.03.03 12:17:00")
}

// Get prefix path
prefix := commandline.PrefixPath()

// Get configuration file
conf := commandline.ConfigFile()
```
