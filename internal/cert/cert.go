package cert

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"github.com/dromara/carbon/v2"
	"math/big"
	"os"
	"path/filepath"
	"strings"
)

const JetProfileCertStr = `-----BEGIN CERTIFICATE-----
MIIFOzCCAyOgAwIBAgIJANJssYOyg3nhMA0GCSqGSIb3DQEBCwUAMBgxFjAUBgNV
BAMMDUpldFByb2ZpbGUgQ0EwHhcNMTUxMDAyMTEwMDU2WhcNNDUxMDI0MTEwMDU2
WjAYMRYwFAYDVQQDDA1KZXRQcm9maWxlIENBMIICIjANBgkqhkiG9w0BAQEFAAOC
Ag8AMIICCgKCAgEA0tQuEA8784NabB1+T2XBhpB+2P1qjewHiSajAV8dfIeWJOYG
y+ShXiuedj8rL8VCdU+yH7Ux/6IvTcT3nwM/E/3rjJIgLnbZNerFm15Eez+XpWBl
m5fDBJhEGhPc89Y31GpTzW0vCLmhJ44XwvYPntWxYISUrqeR3zoUQrCEp1C6mXNX
EpqIGIVbJ6JVa/YI+pwbfuP51o0ZtF2rzvgfPzKtkpYQ7m7KgA8g8ktRXyNrz8bo
iwg7RRPeqs4uL/RK8d2KLpgLqcAB9WDpcEQzPWegbDrFO1F3z4UVNH6hrMfOLGVA
xoiQhNFhZj6RumBXlPS0rmCOCkUkWrDr3l6Z3spUVgoeea+QdX682j6t7JnakaOw
jzwY777SrZoi9mFFpLVhfb4haq4IWyKSHR3/0BlWXgcgI6w6LXm+V+ZgLVDON52F
LcxnfftaBJz2yclEwBohq38rYEpb+28+JBvHJYqcZRaldHYLjjmb8XXvf2MyFeXr
SopYkdzCvzmiEJAewrEbPUaTllogUQmnv7Rv9sZ9jfdJ/cEn8e7GSGjHIbnjV2ZM
Q9vTpWjvsT/cqatbxzdBo/iEg5i9yohOC9aBfpIHPXFw+fEj7VLvktxZY6qThYXR
Rus1WErPgxDzVpNp+4gXovAYOxsZak5oTV74ynv1aQ93HSndGkKUE/qA/JECAwEA
AaOBhzCBhDAdBgNVHQ4EFgQUo562SGdCEjZBvW3gubSgUouX8bMwSAYDVR0jBEEw
P4AUo562SGdCEjZBvW3gubSgUouX8bOhHKQaMBgxFjAUBgNVBAMMDUpldFByb2Zp
bGUgQ0GCCQDSbLGDsoN54TAMBgNVHRMEBTADAQH/MAsGA1UdDwQEAwIBBjANBgkq
hkiG9w0BAQsFAAOCAgEAjrPAZ4xC7sNiSSqh69s3KJD3Ti4etaxcrSnD7r9rJYpK
BMviCKZRKFbLv+iaF5JK5QWuWdlgA37ol7mLeoF7aIA9b60Ag2OpgRICRG79QY7o
uLviF/yRMqm6yno7NYkGLd61e5Huu+BfT459MWG9RVkG/DY0sGfkyTHJS5xrjBV6
hjLG0lf3orwqOlqSNRmhvn9sMzwAP3ILLM5VJC5jNF1zAk0jrqKz64vuA8PLJZlL
S9TZJIYwdesCGfnN2AETvzf3qxLcGTF038zKOHUMnjZuFW1ba/12fDK5GJ4i5y+n
fDWVZVUDYOPUixEZ1cwzmf9Tx3hR8tRjMWQmHixcNC8XEkVfztID5XeHtDeQ+uPk
X+jTDXbRb+77BP6n41briXhm57AwUI3TqqJFvoiFyx5JvVWG3ZqlVaeU/U9e0gxn
8qyR+ZA3BGbtUSDDs8LDnE67URzK+L+q0F2BC758lSPNB2qsJeQ63bYyzf0du3wB
/gb2+xJijAvscU3KgNpkxfGklvJD/oDUIqZQAnNcHe7QEf8iG2WqaMJIyXZlW3me
0rn+cgvxHPt6N4EBh5GgNZR4l0eaFEV+fxVsydOQYo1RIyFMXtafFBqQl6DDxujl
FeU3FZ+Bcp12t7dlM4E0/sS1XdL47CfGVj4Bp+/VbF862HmkAbd7shs7sDQkHbU=
-----END CERTIFICATE-----`

const LicenseServerCertStr = `-----BEGIN CERTIFICATE-----
MIIFTDCCAzSgAwIBAgIJAMCrW9HV+hjZMA0GCSqGSIb3DQEBCwUAMB0xGzAZBgNV
BAMMEkxpY2Vuc2UgU2VydmVycyBDQTAgFw0xNjEwMTIxNDMwNTRaGA8yMTE2MTIy
NzE0MzA1NFowHTEbMBkGA1UEAwwSTGljZW5zZSBTZXJ2ZXJzIENBMIICIjANBgkq
hkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAoT7LvHj3JKK2pgc5f02z+xEiJDcvlBi6
fIwrg/504UaMx3xWXAE5CEPelFty+QPRJnTNnSxqKQQmg2s/5tMJpL9lzGwXaV7a
rrcsEDbzV4el5mIXUnk77Bm/QVv48s63iQqUjVmvjQt9SWG2J7+h6X3ICRvF1sQB
yeat/cO7tkpz1aXXbvbAws7/3dXLTgAZTAmBXWNEZHVUTcwSg2IziYxL8HRFOH0+
GMBhHqa0ySmF1UTnTV4atIXrvjpABsoUvGxw+qOO2qnwe6ENEFWFz1a7pryVOHXg
P+4JyPkI1hdAhAqT2kOKbTHvlXDMUaxAPlriOVw+vaIjIVlNHpBGhqTj1aqfJpLj
qfDFcuqQSI4O1W5tVPRNFrjr74nDwLDZnOF+oSy4E1/WhL85FfP3IeQAIHdswNMJ
y+RdkPZCfXzSUhBKRtiM+yjpIn5RBY+8z+9yeGocoxPf7l0or3YF4GUpud202zgy
Y3sJqEsZksB750M0hx+vMMC9GD5nkzm9BykJS25hZOSsRNhX9InPWYYIi6mFm8QA
2Dnv8wxAwt2tDNgqa0v/N8OxHglPcK/VO9kXrUBtwCIfZigO//N3hqzfRNbTv/ZO
k9lArqGtcu1hSa78U4fuu7lIHi+u5rgXbB6HMVT3g5GQ1L9xxT1xad76k2EGEi3F
9B+tSrvru70CAwEAAaOBjDCBiTAdBgNVHQ4EFgQUpsRiEz+uvh6TsQqurtwXMd4J
8VEwTQYDVR0jBEYwRIAUpsRiEz+uvh6TsQqurtwXMd4J8VGhIaQfMB0xGzAZBgNV
BAMMEkxpY2Vuc2UgU2VydmVycyBDQYIJAMCrW9HV+hjZMAwGA1UdEwQFMAMBAf8w
CwYDVR0PBAQDAgEGMA0GCSqGSIb3DQEBCwUAA4ICAQCJ9+GQWvBS3zsgPB+1PCVc
oG6FY87N6nb3ZgNTHrUMNYdo7FDeol2DSB4wh/6rsP9Z4FqVlpGkckB+QHCvqU+d
rYPe6QWHIb1kE8ftTnwapj/ZaBtF80NWUfYBER/9c6To5moW63O7q6cmKgaGk6zv
St2IhwNdTX0Q5cib9ytE4XROeVwPUn6RdU/+AVqSOspSMc1WQxkPVGRF7HPCoGhd
vqebbYhpahiMWfClEuv1I37gJaRtsoNpx3f/jleoC/vDvXjAznfO497YTf/GgSM2
LCnVtpPQQ2vQbOfTjaBYO2MpibQlYpbkbjkd5ZcO5U5PGrQpPFrWcylz7eUC3c05
UVeygGIthsA/0hMCioYz4UjWTgi9NQLbhVkfmVQ5lCVxTotyBzoubh3FBz+wq2Qt
iElsBrCMR7UwmIu79UYzmLGt3/gBdHxaImrT9SQ8uqzP5eit54LlGbvGekVdAL5l
DFwPcSB1IKauXZvi1DwFGPeemcSAndy+Uoqw5XGRqE6jBxS7XVI7/4BSMDDRBz1u
a+JMGZXS8yyYT+7HdsybfsZLvkVmc9zVSDI7/MjVPdk6h0sLn+vuPC1bIi5edoNy
PdiG2uPH5eDO6INcisyPpLS4yFKliaO4Jjap7yzLU9pbItoWgCAYa2NpxuxHJ0tB
7tlDFnvaRnQukqSG+VqNWg==
-----END CERTIFICATE-----`

var (
	JetProfileCert    *Certificate
	LicenseServerCert *Certificate
)

func init() {
	var err error
	JetProfileCert, err = CreateCertFromPem([]byte(JetProfileCertStr))
	if err != nil {
		panic(err)
	}
	LicenseServerCert, err = CreateCertFromPem([]byte(LicenseServerCertStr))
	if err != nil {
		panic(err)
	}
}

type Certificate struct {
	parent     *Certificate
	cert       *x509.Certificate
	privateKey *rsa.PrivateKey
}

func (c *Certificate) WriteCertToFile(path string) error {
	if c.cert == nil {
		return errors.New("certificate is nil")
	}
	certificatePEM := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: c.cert.Raw,
	}
	bytes := pem.EncodeToMemory(certificatePEM)
	dir := filepath.Dir(path)
	_ = os.MkdirAll(dir, 0750)
	return os.WriteFile(path, bytes, 0666)
}

func (c *Certificate) WritePrivateKeyToFile(path string) error {
	if c.privateKey == nil {
		return errors.New("private key is nil")
	}
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(c.privateKey),
	}
	bytes := pem.EncodeToMemory(privateKeyPEM)
	dir := filepath.Dir(path)
	_ = os.MkdirAll(dir, 0750)
	return os.WriteFile(path, bytes, 0666)
}

func (c *Certificate) Sign(hashAlgo crypto.Hash, content []byte) ([]byte, error) {
	if c.privateKey == nil {
		return nil, errors.New("private key is nil")
	}
	sha := hashAlgo.New()
	sha.Write(content)
	hashed := sha.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, c.privateKey, hashAlgo, hashed)
}

func (c *Certificate) SignBase64(hashAlgo crypto.Hash, content []byte) (string, error) {
	signature, err := c.Sign(hashAlgo, content)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

func (c *Certificate) Verify(hashAlgo crypto.Hash, content, signature []byte) error {
	sha := hashAlgo.New()
	sha.Write(content)
	hashed := sha.Sum(nil)
	return rsa.VerifyPKCS1v15(c.PublicKey(), hashAlgo, hashed, signature)
}

func (c *Certificate) RawTBS() ([]byte, error) {
	if c.cert == nil {
		return nil, errors.New("certificate is nil")
	}
	return c.cert.RawTBSCertificate, nil
}

func (c *Certificate) Raw() ([]byte, error) {
	if c.cert == nil {
		return nil, errors.New("certificate is nil")
	}
	return c.cert.Raw, nil
}

func (c *Certificate) RawBase64() (string, error) {
	raw, err := c.Raw()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(raw), nil
}

func (c *Certificate) CommonName() string {
	if c.cert == nil {
		return ""
	}
	return c.cert.Subject.CommonName
}

func (c *Certificate) Signature() []byte {
	if c.cert == nil {
		return nil
	}
	return c.cert.Signature
}

func (c *Certificate) PublicKey() *rsa.PublicKey {
	if c.privateKey != nil {
		return &c.privateKey.PublicKey
	}
	if c.cert != nil {
		return c.cert.PublicKey.(*rsa.PublicKey)
	}
	return nil
}

func (c *Certificate) PrivateKey() *rsa.PrivateKey {
	if c.privateKey == nil {
		return nil
	}
	return c.privateKey
}

func GenerateFakeCertificate(parentCommonName, commonName, certPath, keyPath string) (*Certificate, error) {
	parentCert, err := GenerateCertificate(parentCommonName, nil)
	if err != nil {
		return nil, err
	}
	cert, err := GenerateCertificate(commonName, parentCert)
	if err != nil {
		return nil, err
	}
	if certPath != "" {
		if err = cert.WriteCertToFile(certPath); err != nil {
			return nil, err
		}
	}
	if keyPath != "" {
		if err = cert.WritePrivateKeyToFile(keyPath); err != nil {
			return nil, err
		}
	}
	return cert, nil
}

func GenerateCertificate(commonName string, parent *Certificate) (*Certificate, error) {
	serialNumber, _ := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	template := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: commonName,
		},
		NotBefore:          carbon.Now().StdTime(),
		NotAfter:           carbon.Now().AddYears(2).StdTime(),
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	var (
		parentTemplate   *x509.Certificate
		parentPrivateKey *rsa.PrivateKey
	)

	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		//template.BasicConstraintsValid = true
		//template.IsCA = true
		parentTemplate = template
		parentPrivateKey = privateKey
	} else {
		parentTemplate = parent.cert
		parentPrivateKey = parent.privateKey
	}
	certBytes, err := x509.CreateCertificate(rand.Reader, template, parentTemplate, &privateKey.PublicKey, parentPrivateKey)
	if err != nil {
		return nil, err
	}
	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, err
	}
	certificate := &Certificate{
		parent:     parent,
		cert:       cert,
		privateKey: privateKey,
	}
	if parent == nil {
		certificate.parent = certificate
	}
	return certificate, nil
}

const PemPrefix = "-----"

func MustCreateCertFromFile(certPath, keyPath string) *Certificate {
	cert, err := CreateCertFromFile(certPath, keyPath)
	if err != nil {
		panic(err)
	}
	return cert
}

func CreateCertFromFile(certPath, keyPath string) (cert *Certificate, err error) {
	if cert, err = CreateCertFromFileWithoutPrivateKey(certPath); err != nil {
		return
	}
	if cert.privateKey, err = CreatePrivateKeyFromFile(keyPath); err != nil {
		return nil, err
	}
	return cert, nil
}

func CreateCertFromFileWithoutPrivateKey(certPath string) (*Certificate, error) {
	certBytes, err := os.ReadFile(certPath)
	if err != nil {
		return nil, err
	}
	var cert *Certificate
	if strings.HasPrefix(string(certBytes[:len(PemPrefix)]), PemPrefix) {
		cert, err = CreateCertFromPem(certBytes)
	} else {
		certBytes, _ = base64.StdEncoding.DecodeString(string(certBytes))
		cert, err = CreateCertFromDer(certBytes)
	}
	if err != nil {
		return nil, err
	}
	return cert, nil
}

func CreatePrivateKeyFromFile(path string) (*rsa.PrivateKey, error) {
	keyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var key *rsa.PrivateKey
	if strings.HasPrefix(string(keyBytes[:len(PemPrefix)]), PemPrefix) {
		key, err = CreatePrivateKeyFromPem(keyBytes)
	} else {
		key, err = CreatePrivateKeyFromDer(keyBytes)
	}
	if err != nil {
		return nil, err
	}
	return key, nil
}

func CreateCertFromPem(bytes []byte) (*Certificate, error) {
	block, _ := pem.Decode(bytes)
	return CreateCertFromDer(block.Bytes)
}

func CreatePrivateKeyFromPem(bytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(bytes)
	return CreatePrivateKeyFromDer(block.Bytes)
}

func CreateCertFromDer(bytes []byte) (*Certificate, error) {
	cert, err := x509.ParseCertificate(bytes)
	if err != nil {
		return nil, err
	}
	return &Certificate{
		cert: cert,
	}, nil
}

func CreatePrivateKeyFromDer(bytes []byte) (*rsa.PrivateKey, error) {
	key, err := x509.ParsePKCS1PrivateKey(bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}
