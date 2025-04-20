package cert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"github.com/dromara/carbon/v2"
	"math/big"
	"os"
	"strings"
)

const JBRootCerStr = `-----BEGIN CERTIFICATE-----
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

var JBRootCer *x509.Certificate

func init() {
	var err error
	JBRootCer, err = CreateCertFromPem([]byte(JBRootCerStr))
	if err != nil {
		panic(err)
	}
}

func GenerateCertificate(ca *x509.Certificate, caPrivateKey, userPrivateKey *rsa.PrivateKey, commonName string, signatureAlgorithm x509.SignatureAlgorithm) ([]byte, error) {
	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, err
	}

	template := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: commonName,
		},
		NotBefore:          carbon.Now().StdTime(),
		NotAfter:           carbon.Now().AddYears(2).StdTime(),
		SignatureAlgorithm: signatureAlgorithm,
	}

	var (
		parent *x509.Certificate
		pub    *rsa.PublicKey
		priv   *rsa.PrivateKey
	)
	if ca == nil || caPrivateKey == nil {
		//template.BasicConstraintsValid = true
		//template.IsCA = true
		parent = template
		pub = &userPrivateKey.PublicKey
		priv = userPrivateKey
	} else {
		parent = ca
		pub = &userPrivateKey.PublicKey
		priv = caPrivateKey
	}
	certificateBytes, err := x509.CreateCertificate(rand.Reader, template, parent, pub, priv)
	if err != nil {
		return nil, err
	}

	return certificateBytes, nil
}

const PemPrefix = "-----"

func MustCreateCertFromFile(path string) *x509.Certificate {
	cert, err := CreateCertFromFile(path)
	if err != nil {
		panic(err)
	}
	return cert
}

func CreateCertFromFile(path string) (*x509.Certificate, error) {
	certBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cert *x509.Certificate
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

func MustCreateRsaKeyFromFile(path string) *rsa.PrivateKey {
	key, err := CreateRsaKeyFromFile(path)
	if err != nil {
		panic(err)
	}
	return key
}

func CreateRsaKeyFromFile(path string) (*rsa.PrivateKey, error) {
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

func CreateCertFromPem(bytes []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(bytes)
	return CreateCertFromDer(block.Bytes)
}

func CreatePrivateKeyFromPem(bytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(bytes)
	return CreatePrivateKeyFromDer(block.Bytes)
}

func CreateCertFromDer(bytes []byte) (*x509.Certificate, error) {
	cert, err := x509.ParseCertificate(bytes)
	if err != nil {
		return nil, err
	}
	return cert, nil
}

func CreatePrivateKeyFromDer(bytes []byte) (*rsa.PrivateKey, error) {
	key, err := x509.ParsePKCS1PrivateKey(bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}
