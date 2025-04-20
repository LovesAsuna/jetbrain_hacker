package handler

import (
	"crypto"
	"encoding/xml"
	"fmt"
	"github.com/LovesAsuna/jetbrains_hacker/internal/cert"
	"net/http"
	"strings"
	"time"
)

type BaseRequest struct {
	Salt      string `form:"salt"`
	UserName  string `form:"userName"`
	MachineId string `form:"machineId"`
}

type Helper struct {
	*CertPool
}

func NewHelper(pool *CertPool) *Helper {
	response := &Helper{
		CertPool: pool,
	}
	return response
}

func (b *Helper) GenerateConfirmationStamp(machineId string) string {
	timeStamp := time.Now().UnixMilli()
	licenseStr := fmt.Sprintf("%d:%s", timeStamp, machineId)
	licenseServerCert := b.LicenseServerCert
	signatureBase64, _ := licenseServerCert.SignBase64(crypto.SHA1, []byte(licenseStr))
	rawUserCertBase64, _ := licenseServerCert.RawBase64()
	return fmt.Sprintf("%s:SHA1withRSA:%s:%s", licenseStr, signatureBase64, rawUserCertBase64)
}

func (b *Helper) GenerateLeaseSignature(serverLease string) string {
	leaseSignature, _ := b.UserCert.SignBase64(crypto.SHA512, []byte(serverLease))
	rawUserCertBase64, _ := b.UserCert.RawBase64()
	return fmt.Sprintf("SHA512withRSA-%s-%s", leaseSignature, rawUserCertBase64)
}

func (b *Helper) GetServerUid() string {
	serverUid := "custom"
	licenseServerCommonName := b.LicenseServerCert.CommonName()
	if strings.Contains(licenseServerCommonName, ".") {
		serverUid = strings.Split(licenseServerCommonName, ".")[0]
	}
	return serverUid
}

func (b *Helper) Sign(content []byte) (result []byte, err error) {
	signAlgo := crypto.SHA1
	signature, err := b.LicenseServerCert.SignBase64(signAlgo, content)
	if err != nil {
		return nil, err
	}
	rawLicenseServerCertBase64, _ := b.LicenseServerCert.RawBase64()
	return []byte(fmt.Sprintf("<!-- SHA1withRSA-%s-%s -->", signature, rawLicenseServerCertBase64)), nil
}

type Signable interface {
	Sign(content []byte) (result []byte, err error)
}

type SignedResponse struct {
	Signable
}

func (t *SignedResponse) Render(w http.ResponseWriter) error {
	t.WriteContentType(w)
	content, err := xml.Marshal(t.Signable)
	if err != nil {
		return err
	}
	signature, err := t.Signable.Sign(content)
	if err != nil {
		return err
	}
	_, err = w.Write(signature)
	if err != nil {
		return err
	}
	_, _ = w.Write([]byte("\n"))
	_, err = w.Write(content)
	return err
}

func (t *SignedResponse) WriteContentType(w http.ResponseWriter) {
	w.Header()["Content-Type"] = []string{"text/xml; charset=utf-8"}
}

const CertPoolKey = "cert_pool"

type CertPool struct {
	LicenseServerCert *cert.Certificate
	UserCert          *cert.Certificate
}
