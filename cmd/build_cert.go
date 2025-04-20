package cmd

import (
	"fmt"
	"github.com/LovesAsuna/jetbrains_hacker/internal/cert"
	"github.com/spf13/cobra"
)

var buildCertCmd = &cobra.Command{
	Use:   "build-cert",
	Short: `Build all needed certificates.`,
	Long:  `Build all needed certificates included user certificate and license server certificate.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := cert.GenerateFakeCertificate(
			cert.JetProfileCert.CommonName(),
			cmd.Flag("user-cert-cn").Value.String(),
			cmd.Flag("user-cert").Value.String(),
			cmd.Flag("user-key").Value.String(),
		); err != nil {
			return err
		}
		if _, err := cert.GenerateFakeCertificate(
			cert.LicenseServerCert.CommonName(),
			fmt.Sprintf("%s.lsrv.jetbrains.com", cmd.Flag("server-uid").Value.String()),
			cmd.Flag("license-server-cert").Value.String(),
			cmd.Flag("license-server-key").Value.String(),
		); err != nil {
			return err
		}
		fmt.Println("build cert successfully!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(buildCertCmd)

	buildCertCmd.Flags().StringP("user-cert-cn", "n", "localhost", "Common name of the user certificate.")
	buildCertCmd.Flags().StringP("server-uid", "s", "custom", "The server uid of license server.")
}
