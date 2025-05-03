package handler

import (
	"github.com/LovesAsuna/jetbrains_hacker/internal/cert"
	"github.com/LovesAsuna/jetbrains_hacker/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Config(context *gin.Context) {
	var (
		configText string
		_type      = context.Param("type")
	)
	switch _type {
	case "dns":
		configText = config.BuildDnsConfig()
	case "url":
		configText = config.BuildUrlConfig()
	default:
		_type = "power"
		pool := context.Value(CertPoolKey).(*CertPool)
		configText = config.BuildPowerConfig(
			[2]*cert.Certificate{
				pool.UserCert, cert.JetProfileCert,
			},
			[2]*cert.Certificate{
				pool.LicenseServerCert, cert.LicenseServerCert,
			},
		)
	}
	context.JSON(
		http.StatusOK,
		struct {
			Type   string `json:"type"`
			Config string `json:"config"`
		}{
			Type:   _type,
			Config: configText,
		},
	)
}
