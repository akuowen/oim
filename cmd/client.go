package cmd

import "github.com/spf13/cobra"
import "github.com/oim/client"

func init() {
	rootCmd.AddCommand(clientCmd)
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client command",
	Run:   clientCmdFunc,
}

// client client函数
func clientCmdFunc(cmd *cobra.Command, args []string) {
	client.RunCui()
}
