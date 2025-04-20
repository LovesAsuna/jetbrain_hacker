package config

import (
	"fmt"
	"github.com/LovesAsuna/jetbrains_hacker/internal/algo"
	"github.com/LovesAsuna/jetbrains_hacker/internal/cert"
	"math/big"
	"strings"
)

func BuildPowerConfig(certs ...[2]*cert.Certificate) string {
	lines := make([]string, 0, 4)
	lines = append(lines, "[Result]")
	for _, certPair := range certs {
		var (
			x, z, r        string
			y              int
			subCert        = certPair[0]
			realParentCert = certPair[1]
		)
		bi := new(big.Int)
		bi.SetBytes(subCert.Signature())
		x = bi.String()
		y = subCert.PublicKey().E
		z = realParentCert.PublicKey().N.String()
		bi = new(big.Int)
		subCertRawTBS, _ := subCert.RawTBS()
		em, err := algo.GetEM(cert.JetProfileCert.PublicKey(), subCertRawTBS)
		if err != nil {
			continue
		}
		bi.SetBytes(em)
		r = bi.String()
		lines = append(lines, fmt.Sprintf("EQUAL,%s,%d,%s->%s", x, y, z, r))
	}
	return strings.Join(lines, "\n")
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
