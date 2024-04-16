package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	ConfigFile string
)

func init() {
	cobra.OnFinalize(initOimConfig)
	rootCmd.PersistentFlags().StringVar(&ConfigFile, "config", "./oim.yaml", "config file (default is ./oim.yaml)")
}

func initOimConfig() {

}

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: oim,
}

func oim(cmd *cobra.Command, args []string) {
	// Do Stuff Here
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
