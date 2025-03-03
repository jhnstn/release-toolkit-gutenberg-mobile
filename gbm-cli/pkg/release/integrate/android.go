package integrate

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/jhnstn/release-toolkit-gutenberg-mobile/gbm-cli/pkg/console"
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/gbm-cli/pkg/gbm"
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/gbm-cli/pkg/release"

	"github.com/jhnstn/release-toolkit-gutenberg-mobile/gbm-cli/pkg/gh"
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/gbm-cli/pkg/repo"
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/gbm-cli/pkg/shell"
)

type AndroidIntegration struct {
}

func (ai AndroidIntegration) UpdateGutenbergConfig(dir string, gbmPr gh.PullRequest) error {
	sp := shell.CmdProps{Dir: dir, Verbose: true}
	git := shell.NewGitCmd(sp)
	prId := gbmPr.Number
	prSha := gbmPr.Head.Sha

	configPath := filepath.Join(dir, "build.gradle")
	config, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`(gutenbergMobileVersion\s*=\s*)'(?:.*)'`)
	if match := re.Match(config); !match {
		return errors.New("cannot find a version in the gradle file")
	}

	var replace string
	if releaseAvailable, err := useRelease(gbmPr.ReleaseVersion); err != nil {
		return fmt.Errorf("unable to check for a release: %s", err)
	} else if releaseAvailable {
		console.Info("Updating gutenberg-mobile ref to the tag v%s", gbmPr.ReleaseVersion)
		replace = fmt.Sprintf(`$1'v%s'`, fmt.Sprint(gbmPr.ReleaseVersion))
	} else {
		console.Info("Updating gutenberg-mobile ref to the commit %s", prSha)
		replace = fmt.Sprintf(`$1'%v-%s'`, prId, prSha)
	}
	config = re.ReplaceAll(config, []byte(replace))

	if err := os.WriteFile(configPath, config, 0644); err != nil {
		return err
	}
	return git.CommitAll("Release script: Update build.gradle gutenbergMobileVersion to ref")
}

func (ai AndroidIntegration) GetRepo() string {
	return repo.WordPressAndroidRepo
}

func (ai AndroidIntegration) GetPr(ri ReleaseIntegration) (gh.PullRequest, error) {
	// @TODO: add support for finding non release PRs
	if ri.Version != "" {
		return release.FindAndroidReleasePr(ri.Version)
	}
	return gh.PullRequest{}, nil
}

func (ai AndroidIntegration) GbPublished(gbmPr gh.PullRequest) (bool, error) {
	published, err := gbm.AndroidGbmBuildPublished(gbmPr)
	if err != nil {
		console.Warn("Error checking if GBM build is published: %v", err)
	}
	return published, nil
}
