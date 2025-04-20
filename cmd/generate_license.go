package cmd

import (
	"crypto"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/LovesAsuna/jetbrains_hacker/internal/cert"
	"github.com/LovesAsuna/jetbrains_hacker/internal/license"
	"github.com/dromara/carbon/v2"
	"github.com/spf13/cobra"
	"math/rand"
	"strings"
)

var generateLicenseCmd = &cobra.Command{
	Use:   "generate-license",
	Short: `generate-license.`,
	Long:  `generate-license.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		licenseId := cmd.Flag("licenseId").Value.String()
		license, err := license.Generate(
			licenseId,
			cmd.Flag("name").Value.String(),
			cmd.Flag("user").Value.String(),
			cmd.Flag("email").Value.String(),
			cmd.Flag("time").Value.String(),
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

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GetRandomString(length int) string {
	if length <= 0 {
		return ""
	}
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(bytes)
}

func init() {
	rootCmd.AddCommand(generateLicenseCmd)

	generateLicenseCmd.Flags().String("licenseId", GetRandomString(10), "Id of license.")
	generateLicenseCmd.Flags().String("name", "user", "The licensee name of license.")
	generateLicenseCmd.Flags().String("user", "user", "The assignee name of license.")
	generateLicenseCmd.Flags().String("email", "i@user.com", "The assignee email of license.")
	generateLicenseCmd.Flags().String("time", carbon.Now().AddYears(2).SetLayout(carbon.DateLayout).String(), "The expire time of license.")
}
