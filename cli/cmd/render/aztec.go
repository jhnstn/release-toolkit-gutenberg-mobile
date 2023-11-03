package render

import (
	"github.com/jhnstn/release-toolkit-gutenberg-mobile/cli/pkg/console"
	"github.com/spf13/cobra"
)

var AztecCmd = &cobra.Command{
	Use:   "aztec",
	Short: "Render the steps for upgrading Aztec",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := renderAztecSteps(false)
		exitIfError(err, 1)

		if writeToClipboard {
			console.Clipboard(result)
		} else {
			console.Out(result)
		}
	},
}
