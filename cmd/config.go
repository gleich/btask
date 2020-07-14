package cmd

import (
	"fmt"

	"github.com/Matt-Gleich/btask/doc"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "⚙️ Configure btask",
	Long: doc.CreateLong(
		`
____ ____ _  _ ____ _ ____
|    |  | |\ | |___ | | __
|___ |__| | \| |    | |__]
`,
	),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
