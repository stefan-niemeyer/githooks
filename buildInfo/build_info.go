package buildInfo

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"time"
)

var (
	version   = "local build"
	gitCommit = ""
	buildDate = fmt.Sprintf("%s", time.Now().UTC().Format(time.RFC3339))
	goos      = runtime.GOOS
	goarch    = runtime.GOARCH
)

// BuildInfo holds information about the build of an executable.
type BuildInfo struct {
	Version   string `json:"version,omitempty"`
	GitCommit string `json:"gitCommit,omitempty"`
	BuildDate string `json:"buildDate,omitempty"`
	GoOs      string `json:"goOs,omitempty"`
	GoArch    string `json:"goArch,omitempty"`
}

// GetBuildInfo returns an instance of BuildInfo.
func GetBuildInfo() BuildInfo {
	return BuildInfo{
		version,
		gitCommit,
		buildDate,
		goos,
		goarch,
	}
}

func (v BuildInfo) ToString() string {
	j, err := json.Marshal(v)
	if err != nil {
		log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	}
	return fmt.Sprintf("%+v", string(j))
}
