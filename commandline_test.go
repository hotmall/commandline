package commandline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmdline(t *testing.T) {
	assert.Equal(t, "/opt/oss/hot", PrefixPath())
	assert.Equal(t, "/opt/oss/hot/etc/conf/hot.json", ConfigFile())
	assert.Equal(t, "/opt/oss/hot/var/log", LogPath())
	assert.Equal(t, 32018, Port())
}
