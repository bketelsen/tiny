package cmd

import (
	"fmt"

	ver "github.com/hashicorp/go-version"
)

var (
	// populated by ldflags
	GitCommit string
	GitTag    string
	BuildDate string

	version    = "v0.1.0"
	prerelease = "" // blank if full release
)

func buildVersion() string {
	verStr := version
	if prerelease != "" {
		verStr = fmt.Sprintf("%s-%s", version, prerelease)
	}

	// check for git tag via ldflags
	if len(GitTag) > 0 {
		verStr = GitTag
	}

	// make sure we fail fast (panic) if bad version - this will get caught in CI tests
	ver.Must(ver.NewVersion(verStr))
	return verStr
}

func buildVersionEnhanced() string {
	verStr := buildVersion()
	if len(description) > 0 {
		verStr = fmt.Sprintf("%s\n%s", verStr, description)
	}
	if len(repository) > 0 {
		verStr = fmt.Sprintf("%s\n\nRepo: %s", verStr, repository)
	}

	if len(docs) > 0 {
		verStr = fmt.Sprintf("%s\nDocs: %s", verStr, docs)
	}
	return verStr
}
