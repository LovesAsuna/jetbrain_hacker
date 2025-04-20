package handler

import (
	"github.com/LovesAsuna/jetbrains_hacker/internal/cert"
	"github.com/LovesAsuna/jetbrains_hacker/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Config(context *gin.Context) {
	var html string
	switch context.Param("type") {
	case "dns":
		html = config.BuildDnsConfig()
	case "url":
		html = config.BuildUrlConfig()
	default:
		pool := context.Value(CertPoolKey).(*CertPool)
		html = config.BuildPowerConfig(
			[2]*cert.Certificate{
				pool.UserCert, cert.JetProfileCert,
			},
			[2]*cert.Certificate{
				pool.LicenseServerCert, cert.LicenseServerCert,
			},
		)
	}
	context.HTML(http.StatusOK, "config.html", map[string]string{"content": html})
}
