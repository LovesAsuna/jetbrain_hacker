package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "jetbrain-hacker",
	Short: "Generate custom license code or run a license server.",
	Long:  `Generate custom license code or run a license server.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("user-cert", "c", "cert/user.crt", "Path to store the user certificate.")
	rootCmd.PersistentFlags().StringP("user-key", "k", "cert/user.key", "Path to store the user private key.")
	rootCmd.PersistentFlags().String("license-server-cert", "cert/license_server.crt", "Path to store the license server certificate.")
	rootCmd.PersistentFlags().String("license-server-key", "cert/license_server.key", "Path to store the license server private key.")
}
