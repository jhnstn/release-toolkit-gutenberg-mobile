package cmd

import (
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/cli/cmd/release"
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/cli/cmd/render"
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/cli/pkg/console"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gbm",
	Short: "Gutenberg Mobile CLI",
}

func Execute() {
	err := rootCmd.Execute()
	console.ExitIfError(err)
}

func init() {
	// Add the render command
	rootCmd.AddCommand(render.RenderCmd)
	rootCmd.AddCommand(release.ReleaseCmd)
}
