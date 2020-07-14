package cmd

import (
	"github.com/Matt-Gleich/statuser"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "btask",
	Short: "✅ A beautiful todo list for your terminal",
	Long: `
  ___  ___ ____ ____ _  _
  |__]  |  |__| [__  |_/
  |__]  |  |  | ___] | \_
  ⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯

  🐙 GitHub: https://github.com/Matt-Gleich/btask
  📜 Produced under the MIT license
  🛠  Authors:
	- Matthew Gleich (https://github.com/Matt-Gleich)

`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		statuser.Error("Failed to execute root command", err, 1)
	}
}
