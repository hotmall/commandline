package commandline

import (
	"fmt"
)

var (
	//Version define software version
	Version string
	// CommitHash represents the Git commit hash at built time
	CommitHash string
	// BuildDate represents the date when this tool was built
	BuildDate string
	// GoVersion represents go version
	GoVersion string
)

// ShowVersion print version infomation
func ShowVersion() {
	fmt.Printf("Version\t\t: v%v\nCommit Hash\t: %v\nBuild Date\t: %v\n\n%v\n",
		Version, CommitHash, BuildDate, GoVersion)
}
