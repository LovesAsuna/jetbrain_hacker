package cmd

import (
	"context"
	"crypto"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/LovesAsuna/jetbrains_hacker/internal/cert"
	"github.com/LovesAsuna/jetbrains_hacker/internal/license"
	"github.com/LovesAsuna/jetbrains_hacker/internal/util"
	"github.com/dromara/carbon/v2"
	"github.com/lovesasuna/sync/coroutinegroup"
	"github.com/spf13/cobra"
	"strings"
)

var generateLicenseCmd = &cobra.Command{
	Use:   "generate-license",
	Short: `generate-license.`,
	Long:  `generate-license.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		licenseId := cmd.Flag("licenseId").Value.String()
		codes, err := getCodes()
		if err != nil {
			return err
		}
		license, err := license.Generate(
			licenseId,
			cmd.Flag("name").Value.String(),
			cmd.Flag("user").Value.String(),
			cmd.Flag("email").Value.String(),
			cmd.Flag("time").Value.String(),
			codes...,
		)
		if err != nil {
			return err
		}
		licenseJs, _ := json.Marshal(license)
		licensePartBase64 := base64.StdEncoding.EncodeToString(licenseJs)

		userCert := cert.MustCreateCertFromFile(cmd.Flag("user-cert").Value.String(), cmd.Flag("user-key").Value.String())
		certPartBase64, _ := userCert.RawBase64()

		signatureBytes, _ := userCert.Sign(crypto.SHA1, licenseJs)
		signatureBase64 := base64.StdEncoding.EncodeToString(signatureBytes)

		l := strings.Join([]string{licenseId, licensePartBase64, signatureBase64, certPartBase64}, "-")
		fmt.Println(l)
		return nil
	},
}

func getCodes() ([]string, error) {
	var (
		productCodes []string
		pluginCodes  []string
	)
	group, _ := coroutinegroup.WithContext(context.Background())
	group.Go(
		func(ctx context.Context) error {
			codes, err := license.GetProductCode()
			if err != nil {
				return err
			}
			productCodes = codes
			return nil
		},
	)
	group.Go(
		func(ctx context.Context) error {
			codes, err := license.GetPluginCode(10000, 0, "")
			if err != nil {
				return err
			}
			pluginCodes = codes
			return nil
		},
	)
	errs := group.Wait()
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}
	codes := make([]string, 0, len(productCodes)+len(pluginCodes))
	codes = append(codes, productCodes...)
	codes = append(codes, pluginCodes...)
	return codes, nil
}

func init() {
	rootCmd.AddCommand(generateLicenseCmd)

	generateLicenseCmd.Flags().String("licenseId", util.GetRandomString(10), "Id of license.")
	generateLicenseCmd.Flags().String("name", "user", "The licensee name of license.")
	generateLicenseCmd.Flags().String("user", "user", "The assignee name of license.")
	generateLicenseCmd.Flags().String("email", "i@user.com", "The assignee email of license.")
	generateLicenseCmd.Flags().String("time", carbon.Now().AddYears(2).SetLayout(carbon.DateLayout).String(), "The expire time of license.")
}
