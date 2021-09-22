package cmd

import (
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the status of Colima",
	Long:  `Show the status of Colima`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(app.Status())
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
