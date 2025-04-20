package cmd

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/LovesAsuna/jetbrains_hacker/internal/algo"
	"github.com/LovesAsuna/jetbrains_hacker/internal/cert"
	"github.com/spf13/cobra"
	"os"
)

var buildCertCmd = &cobra.Command{
	Use:   "build-cert",
	Short: `Build all needed certificates.`,
	Long:  `Build all needed certificates included fake root certificate and user certificate.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		caParam := &BuildCertParam{
			SignatureAlgorithm: x509.SHA256WithRSA,
			CommonName:         cert.JBRootCer.Subject.CommonName,
			ResultCertFile:     cmd.Flag("fake-root-cert").Value.String(),
			ResultKeyFile:      cmd.Flag("fake-root-key").Value.String(),
			IsCA:               true,
		}
		if b, _ := cmd.Flags().GetBool("keep-fake-root-cert"); !b {
			caParam.ResultCertFile = ""
			caParam.ResultKeyFile = ""
		}
		fakeRootCert, fakeRootPrivateKey, err := doBuildCert(caParam)
		if err != nil {
			return err
		}
		userParam := &BuildCertParam{
			SignatureAlgorithm: x509.SHA256WithRSA,
			CommonName:         cmd.Flag("user-cert-cn").Value.String(),
			Ca:                 fakeRootCert,
			CaPrivateKey:       fakeRootPrivateKey,
			ResultCertFile:     cmd.Flag("user-cert").Value.String(),
			ResultKeyFile:      cmd.Flag("user-key").Value.String(),
			IsCA:               false,
		}
		_, _, err = doBuildCert(userParam)
		if err != nil {
			return err
		}
		fmt.Println("root and user cert build successfully!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(buildCertCmd)

	buildCertCmd.Flags().String("fake-root-cert", "./fake.cer", "Path to store the fake root cer.")
	buildCertCmd.Flags().String("fake-root-key", "./fake.key", "Path to store the fake root private key.")
	buildCertCmd.Flags().Bool("keep-fake-root-cert", false, "Keep fake root cert after building user cert.")
	buildCertCmd.Flags().StringP("user-cert", "c", "./user.cer", "Path to store the user cer.")
	buildCertCmd.Flags().StringP("user-key", "k", "./user.key", "Path to store the user private key.")
	buildCertCmd.Flags().StringP("user-cert-cn", "n", "LovesAsuna", "Common name of user certificate.")
}

type BuildCertParam struct {
	SignatureAlgorithm x509.SignatureAlgorithm
	CommonName         string
	ResultCertFile     string
	ResultKeyFile      string
	Ca                 *x509.Certificate
	CaPrivateKey       *rsa.PrivateKey
	IsCA               bool
}

func doBuildCert(param *BuildCertParam) (*x509.Certificate, *rsa.PrivateKey, error) {
	privateKey, err := algo.GenerateKeyPair(4096)
	if err != nil {
		return nil, nil, err
	}

	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	if param.ResultCertFile != "" {
		privateKeyFile, err := os.Create(param.ResultKeyFile)
		if err != nil {
			return nil, nil, err
		}
		defer privateKeyFile.Close()
		_ = pem.Encode(privateKeyFile, privateKeyPEM)
	}

	var certificateBytes []byte
	if param.IsCA {
		certificateBytes, err = cert.GenerateCertificate(nil, nil, privateKey, param.CommonName, param.SignatureAlgorithm)
	} else {
		certificateBytes, err = cert.GenerateCertificate(param.Ca, param.CaPrivateKey, privateKey, param.CommonName, param.SignatureAlgorithm)
	}
	if err != nil {
		return nil, nil, err
	}

	certificatePEM := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certificateBytes,
	}

	if param.ResultCertFile != "" {
		certificateFile, err := os.Create(param.ResultCertFile)
		if err != nil {
			return nil, nil, err
		}
		defer certificateFile.Close()
		_ = pem.Encode(certificateFile, certificatePEM)
	}

	cert, err := x509.ParseCertificate(certificateBytes)
	if err != nil {
		panic(err)
	}
	return cert, privateKey, nil
}
