package cmd

import (
	"fmt"
	"github.com/LovesAsuna/jetbrains_hacker/internal/cert"
	"github.com/LovesAsuna/jetbrains_hacker/internal/config"
	"github.com/spf13/cobra"
)

var buildConfigCmd = &cobra.Command{
	Use:   "build-config",
	Short: `Build the *.conf of ja-netfilter.`,
	Long:  `Build the *.conf of ja-netfilter.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		switch cmd.Flag("type").Value.String() {
		case "dns":
			fmt.Println(config.BuildDnsConfig())
			return nil
		case "url":
			fmt.Println(config.BuildUrlConfig())
			return nil
		case "power":
			userCert, err := cert.CreateCertFromFileWithoutPrivateKey(cmd.Flag("user-cert").Value.String())
			if err != nil {
				return err
			}
			licenseServerCert, err := cert.CreateCertFromFileWithoutPrivateKey(cmd.Flag("license-server-cert").Value.String())
			if err != nil {
				return err
			}
			fmt.Println(
				config.BuildPowerConfig(
					[2]*cert.Certificate{
						userCert, cert.JetProfileCert,
					},
					[2]*cert.Certificate{
						licenseServerCert, cert.LicenseServerCert,
					},
				),
			)
		default:
			fmt.Println("unknown config type.")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(buildConfigCmd)

	buildConfigCmd.Flags().StringP("type", "t", "power", "If empty use power. Possible values: 'power', 'dns', 'url'.")
}
