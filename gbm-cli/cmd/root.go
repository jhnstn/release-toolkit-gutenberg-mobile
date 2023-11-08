package cmd

import (
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/gbm-cli/cmd/release"
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/gbm-cli/cmd/render"
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/gbm-cli/pkg/console"
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/gbm-cli/pkg/gh"
	"github.com/spf13/cobra"
)

const Version = "v1.2.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gbm-cli",
	Short:   "Gutenberg Mobile CLI",
	Version: Version,
}

func Execute() {
	err := rootCmd.Execute()
	console.ExitIfError(err)

}

func init() {
	// Add the render command
	rootCmd.AddCommand(render.RenderCmd)
	rootCmd.AddCommand(release.ReleaseCmd)

	// Check to see if the user is running the latest version
	// of the CLI. If not, let them know.
	latestRelease, err := gh.GetLatestRelease("release-toolkit-gutenberg-mobile")
	console.ExitIfError(err)

	if latestRelease.TagName != Version {
		console.Warn("You are running an older version of the CLI. Please update to %s", latestRelease.TagName)
	}
}
