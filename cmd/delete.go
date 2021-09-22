package cmd

import (
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete and teardown Colima",
	Long: `Delete and teardown Colima and all settings.

Use with caution. This deletes everything and a startup afterwards is like the
initial startup of Colima.

If you simply want to reset the Kubernetes cluster, run 'colima kubernetes reset'.`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(app.Delete())
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
