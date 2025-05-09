package license

import (
	"crypto"
	"encoding/base64"
	"encoding/json"
	"github.com/LovesAsuna/jetbrains_hacker/internal/cert"
	"github.com/dromara/carbon/v2"
	"strings"
)

type License struct {
	LicenseID          string    `json:"licenseId"`
	LicenseeName       string    `json:"licenseeName"`
	AssigneeName       string    `json:"assigneeName"`
	AssigneeEmail      string    `json:"assigneeEmail"`
	LicenseRestriction string    `json:"licenseRestriction"`
	CheckConcurrentUse bool      `json:"checkConcurrentUse"`
	Products           []Product `json:"products"`
	Metadata           string    `json:"metadata"`
	Hash               string    `json:"hash"`
	GracePeriodDays    int       `json:"gracePeriodDays"`
	AutoProlongated    bool      `json:"autoProlongated"`
	IsAutoProlongated  bool      `json:"isAutoProlongated"`
}

type Product struct {
	Code         string `json:"code"`
	FallbackDate string `json:"fallbackDate"`
	PaidUpTo     string `json:"paidUpTo"`
}

func GenerateLicenseCode(cert *cert.Certificate, licenseId, licenseeName, assigneeName, assigneeEmail, time string, codes ...string) (string, error) {
	license, err := GenerateLicense(
		licenseId,
		licenseeName,
		assigneeName,
		assigneeEmail,
		time,
		codes...,
	)
	if err != nil {
		return "", err
	}
	licenseJs, _ := json.Marshal(license)
	licensePartBase64 := base64.StdEncoding.EncodeToString(licenseJs)

	certPartBase64, _ := cert.RawBase64()

	signatureBytes, _ := cert.Sign(crypto.SHA1, licenseJs)
	signatureBase64 := base64.StdEncoding.EncodeToString(signatureBytes)

	return strings.Join([]string{licenseId, licensePartBase64, signatureBase64, certPartBase64}, "-"), nil
}

func GenerateLicense(licenseId, licenseeName, assigneeName, assigneeEmail, time string, codes ...string) (*License, error) {
	fallBackDate := carbon.Now().SetLayout(carbon.DateLayout).String()
	paidUpTo := carbon.ParseByLayout(time, carbon.DateLayout).String()

	products := make([]Product, 0, len(codes))
	for _, code := range codes {
		products = append(products, Product{
			Code:         code,
			FallbackDate: fallBackDate,
			PaidUpTo:     paidUpTo,
		})
	}

	license := &License{
		LicenseID:         licenseId,
		LicenseeName:      licenseeName,
		AssigneeName:      assigneeName,
		AssigneeEmail:     assigneeEmail,
		Products:          products,
		Metadata:          "0120230102PPAA013009",
		Hash:              "41472961/0:1563609451",
		GracePeriodDays:   7,
		AutoProlongated:   true,
		IsAutoProlongated: true,
	}
	return license, nil
}
