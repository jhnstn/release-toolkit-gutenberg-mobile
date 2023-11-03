package prepare

import (
	"errors"

	"github.com/jhnstn/release-toolkit-gutenberg-mobile/cli/cmd/utils"
	wp "github.com/jhnstn/release-toolkit-gutenberg-mobile/cli/cmd/workspace"
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/cli/pkg/console"
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/cli/pkg/gbm"
	"github.com/spf13/cobra"
)

var exitIfError func(error, int)
var keepTempDir, noTag bool
var workspace wp.Workspace
var tempDir, version string

var PrepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "Prepare for a release",
}

func Execute() {
	err := PrepareCmd.Execute()
	exitIfError(err, 1)
}

// Set up the temp directory and version
// Also validate Aztec versions
func preflight(args []string) {
	var err error
	tempDir = workspace.Dir()

	semver, err := utils.GetVersionArg(args)
	exitIfError(err, 1)
	version = semver.String()

	// Validate Aztec version
	if valid := gbm.ValidateAztecVersions(); !valid {
		exitIfError(errors.New("invalid Aztec versions found"), 1)
	}
	if keepTempDir {
		workspace.Keep()
	}
}

func init() {
	var err error
	workspace, err = wp.NewWorkspace()
	utils.ExitIfError(err, 1)

	exitIfError = func(err error, code int) {
		if err != nil {
			console.Error(err)
			utils.Exit(code, workspace.Cleanup)
		}
	}

	PrepareCmd.AddCommand(gbmCmd)
	PrepareCmd.AddCommand(gbCmd)
	PrepareCmd.AddCommand(allCmd)
	PrepareCmd.PersistentFlags().BoolVar(&keepTempDir, "k", false, "Keep temporary directory after running command")
	PrepareCmd.PersistentFlags().BoolVar(&noTag, "no-tag", false, "Prevent tagging the release")

}
