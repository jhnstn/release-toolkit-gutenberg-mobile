package release

import (
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/cli/pkg/gh"
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/cli/pkg/semver"
)

type Build struct {
	Version semver.SemVer
	Dir     string
	UseTag  bool
	Repo    string
	Prs     []gh.PullRequest
	Base    gh.Repo
	Depth   string
}

type ReleaseChanges struct {
	Title  string
	Number int
	PrUrl  string
	Issues []string
}
