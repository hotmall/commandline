package commandline

import "testing"

func TestVersion(t *testing.T) {
	Version = "1.0.2"
	CommitHash = "fc087a8"
	BuildDate = "2019-03-11T20:52:45+0800"

	ShowVersion()
}
