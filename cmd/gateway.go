package cmd

import (
	"github.com/oim/gateway"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(gatewayCmd)
}

var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "gateway command",
	Run:   gatewayCmdFunc,
}

func gatewayCmdFunc(cmd *cobra.Command, args []string) {
	gateway.RunGateway(ConfigFile)
}
