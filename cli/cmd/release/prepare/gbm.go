package prepare

import (
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/cli/pkg/console"
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/cli/pkg/release"
	"github.com/spf13/cobra"
)

var gbmCmd = &cobra.Command{
	Use:   "gbm",
	Short: "prepare Gutenberg Mobile release",
	Long:  `Use this command to prepare a Gutenberg Mobile release PR`,
	Run: func(cmd *cobra.Command, args []string) {
		preflight(args)
		defer workspace.Cleanup()

		console.Info("Preparing Gutenberg Mobile for release %s", version)

		pr, err := release.CreateGbmPR(version, tempDir)
		exitIfError(err, 1)

		console.Info("Created PR %s", pr.Url)
	},
}
