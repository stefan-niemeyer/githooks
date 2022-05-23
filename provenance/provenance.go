// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package provenance

import (
	"fmt"
	"runtime"
)

var (
	version = "v0.7.0"
	goos    = runtime.GOOS
	goarch  = runtime.GOARCH
)

// Provenance holds information about the build of an executable.
type Provenance struct {
	// Version of the kustomize binary.
	Version string `json:"version,omitempty"`
	// GitCommit is a git commit
	GoOs string `json:"goOs,omitempty"`
	// GoArch holds architecture name.
	GoArch string `json:"goArch,omitempty"`
}

// GetProvenance returns an instance of Provenance.
func GetProvenance() Provenance {
	return Provenance{
		version,
		goos,
		goarch,
	}
}

// Full returns the full provenance stamp.
func (v Provenance) Full() string {
	return fmt.Sprintf("%+v", v)
}
