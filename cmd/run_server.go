package cmd

import (
	"github.com/LovesAsuna/jetbrains_hacker/server"
	"github.com/LovesAsuna/jetbrains_hacker/server/config"
	"github.com/spf13/cobra"
)

var runServerCmd = &cobra.Command{
	Use:   "run-server",
	Short: `Run a JetBrain license server.`,
	Long:  `Run a JetBrain license server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config.InitServerConfig(
			&config.ServerConfig{
				Addr:                        cmd.Flag("addr").Value.String(),
				Licensee:                    cmd.Flag("licensee").Value.String(),
				UserCertPath:                cmd.Flag("user-cert").Value.String(),
				UserPrivateKeyPath:          cmd.Flag("user-key").Value.String(),
				LicenseServerCertPath:       cmd.Flag("license-server-cert").Value.String(),
				LicenseServerPrivateKeyPath: cmd.Flag("license-server-key").Value.String(),
			},
		)
		return server.RunServer()
	},
}

func init() {
	rootCmd.AddCommand(runServerCmd)

	runServerCmd.Flags().String("addr", ":80", "The address of license server.")
	runServerCmd.Flags().String("licensee", "", "The licensee of license server. Default to computer user name")
}
