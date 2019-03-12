package commandline

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func pidFile() string {
	logPath := LogPath()
	if err := os.MkdirAll(logPath, os.ModeDir); err != nil {
		log.Fatal(err)
	}
	pidFile := filepath.Base(os.Args[0])
	pidFile = strings.TrimSuffix(pidFile, path.Ext(pidFile))
	pidFile = logPath + "/" + pidFile + ".pid"
	return pidFile
}

// WriteProcID write procID into .pid file
func WriteProcID() {
	pidFile := pidFile()
	pid := strconv.Itoa(os.Getpid())
	err := ioutil.WriteFile(pidFile, []byte(pid), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// ReadProcID read procID from .pid file
func ReadProcID() int {
	pidFile := pidFile()

	content, err := ioutil.ReadFile(pidFile)
	if err != nil {
		log.Fatal(err)
	}

	procID, err := strconv.Atoi(string(content))
	if err != nil {
		log.Fatal(err)
	}
	return procID
}

// Exit send a SIGINT(2) signal to the process
func Exit() {
	pid := ReadProcID()
	p, err := os.FindProcess(pid)
	if err != nil {
		log.Fatal(err)
	}

	p.Signal(os.Interrupt)
}
