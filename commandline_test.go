package commandline_test

import (
	"strings"
	"testing"

	"github.com/mallbook/commandline"
)

func TestCmdline(t *testing.T) {
	prefix := commandline.PrefixPath()
	if strings.Compare(prefix, "/opt/oss/hot/myapp") != 0 {
		t.Errorf("The prefix path expect /opt/oss/MBK/myapp, but %s\n", prefix)
	}

	config := commandline.ConfigFile()
	if strings.Compare(config, "etc/conf/hot.json") != 0 {
		t.Errorf("The config expect etc/conf/myapp.json, but %s\n", config)
	}

	port := commandline.Port()
	if port != 32019 {
		t.Errorf("The Port except 32019, But %d\n", port)
	}

	logPath := commandline.LogPath()
	if strings.Compare(logPath, "var/log/myapp") != 0 {
		t.Errorf("The logpath expect var/log/myapp, but %s\n", logPath)
	}
}
