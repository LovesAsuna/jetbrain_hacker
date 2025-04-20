package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/LovesAsuna/jetbrains_hacker/internal/algo"
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

		priv := cert.MustCreateRsaKeyFromFile(cmd.Flag("user-key").Value.String())
		signatureBytes := algo.Sign(licenseJs, priv)
		signatureBase64 := base64.StdEncoding.EncodeToString(signatureBytes)

		userCert := cert.MustCreateCertFromFile(cmd.Flag("user-cert").Value.String())
		certPartBase64 := base64.StdEncoding.EncodeToString(userCert.Raw) // 4

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

	generateLicenseCmd.Flags().String("user-cert", "./user.cer", "Path to store the user cer.")
	generateLicenseCmd.Flags().String("user-key", "./user.key", "Path to store the user private key.")
	generateLicenseCmd.Flags().String("licenseId", GetRandomString(10), "Id of license.")
	generateLicenseCmd.Flags().String("name", "LovesAsuna", "The licensee name of license.")
	generateLicenseCmd.Flags().String("user", "LovesAsuna", "The assignee name of license.")
	generateLicenseCmd.Flags().String("email", "qq625924077@gmail.com", "The assignee email of license.")
	generateLicenseCmd.Flags().String("time", carbon.Now().AddYears(2).SetLayout(carbon.DateLayout).String(), "The expire time of license.")
}
