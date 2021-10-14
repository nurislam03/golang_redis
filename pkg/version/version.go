package version

import (
	"runtime"
	"strings"

	"github.com/nurislam03/golang_redis/pkg/helpmate"
)

var (
	// version is the current version of Helm.
	// Update this whenever making a new release.
	// The version is of the format Major.Minor.Patch[-Prerelease][+BuildMetadata]
	//
	// Increment major number for new feature additions and behavioral changes.
	// Increment minor number for bug fixes and performance enhancements.
	version = "v0.0"

	// metadata is extra build time data
	metadata = ""
	// gitCommit is the git sha1
	gitCommit = ""
	// GitBranch is the git current branch name.
	gitBranch = ""
	// gitTreeState is the state of the git tree
	gitTreeState = ""
)

// BuildInfo describes the compile time information.
type BuildInfo struct {
	// Version is the current semver.
	Version string `json:"version,omitempty"`
	// GitCommit is the git sha1.
	GitCommit string `json:"git_commit,omitempty"`
	// GitBranch is the git current branch name.
	GitBranch string `json:"git_branch,omitempty"`
	// GitTreeState is the state of the git tree.
	GitTreeState string `json:"git_tree_state,omitempty"`
	// GoVersion is the version of the Go compiler used.
	GoVersion string `json:"go_version,omitempty"`
}

// GetVersion returns the semver string of the version
func GetVersion() string {
	if metadata == "" {
		return version
	}
	return version + "+" + metadata
}

// GetUserAgent returns a user agent for user with an HTTP client
func GetUserAgent() string {
	return "Settlement/" + strings.TrimPrefix(GetVersion(), "v")
}

// Get returns build info
func Get() BuildInfo {
	v := BuildInfo{
		Version:      GetVersion(),
		GitCommit:    gitCommit,
		GitBranch:    gitBranch,
		GitTreeState: gitTreeState,
		GoVersion:    runtime.Version(),
	}

	if helpmate.IsTestEnv() {
		v.GoVersion = ""
	}
	return v
}
