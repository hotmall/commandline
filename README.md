# commandline

## Installation
You can use `go get` to get the latest version:
```sh
go get -u github.com/mallbook/commandline
```

## Quick Start
### Show Help
```
Usage: commandline-test [-hv] [-c file] [-p prefix]

Options:
  -c file
    	set configuration file (default "etc/conf/nginx.conf")
  -h	this help
  -p prefix
    	set prefix path (default ".")
  -v	show version and exit
```

### Usage
```go
// Print Version
if commandline.IsShowVersion() {
    fmt.Println("version: 1.0.1, build: 2019.03.03 12:17:00")
}

// Get prefix path
prefix := commandline.GetPrefixPath()

// Get configuration file
conf := commandline.GetConfigFile()
```
