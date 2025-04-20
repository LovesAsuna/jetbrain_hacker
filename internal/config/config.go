package config

import (
	"crypto/rsa"
	"crypto/x509"
	"github.com/LovesAsuna/jetbrains_hacker/internal/algo"
	"math/big"
	"strings"
)

func BuildPowerConfig(userCert, rootCert *x509.Certificate) string {
	builder := new(strings.Builder)
	builder.WriteString("[Result]\n")
	builder.WriteString("EQUAL,")
	bi := new(big.Int)
	bi.SetBytes(userCert.Signature)
	builder.WriteString(bi.String())
	builder.WriteString(",65537,")
	bi = new(big.Int)
	bi.SetBytes(rootCert.PublicKey.(*rsa.PublicKey).N.Bytes())
	builder.WriteString(bi.String())
	builder.WriteString("->")
	bi = new(big.Int)
	em, err := algo.GetEM(rootCert.PublicKey.(*rsa.PublicKey), userCert.RawTBSCertificate)
	if err != nil {
		panic(err)
	}
	bi.SetBytes(em)
	builder.WriteString(bi.String())
	builder.WriteString("\n\n[Args]")
	return builder.String()
}

func BuildDnsConfig() string {
	builder := new(strings.Builder)
	builder.WriteString("[DNS]\n")
	builder.WriteString("EQUAL,jetbrains.com\n")
	builder.WriteString("EQUAL,plugin.obroom.com")
	return builder.String()
}

func BuildUrlConfig() string {
	builder := new(strings.Builder)
	builder.WriteString("[URL]\n")
	builder.WriteString("PREFIX,https://account.jetbrains.com/lservice/rpc/validateKey.action\n")
	builder.WriteString("PREFIX,https://account.jetbrains.com.cn/lservice/rpc/validateKey.action")
	return builder.String()
}
