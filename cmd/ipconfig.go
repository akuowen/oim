package cmd

import (
	"github.com/oim/ipconfig"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ipConfigCmd)
}

var ipConfigCmd = &cobra.Command{

	Use:   "ipConfig",
	Short: "ipConfig command",
	Run:   ipConfigCmdFunc,
}

// client client函数
func ipConfigCmdFunc(cmd *cobra.Command, args []string) {
	ipconfig.RunIpConfigServer(ConfigFile)
}
