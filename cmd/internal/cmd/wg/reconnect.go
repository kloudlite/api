package wg

import (
	"github.com/spf13/cobra"
)

var reconnectCmd = &cobra.Command{
	Use:   "reconnect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopService()
		startService()
	},
}
