package handler

import (
	"github.com/LovesAsuna/jetbrains_hacker/internal/license"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func License(context *gin.Context) {
	pool := context.Value(CertPoolKey).(*CertPool)
	type Param struct {
		LicenseId string `form:"licenseId"`
		Name      string `form:"name"`
		User      string `form:"user"`
		Email     string `form:"email"`
		Time      string `form:"time"`
		Codes     string `form:"codes"`
	}
	param := new(Param)
	err := context.Bind(param)
	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	licenseCode, err := license.GenerateLicenseCode(
		pool.UserCert,
		param.LicenseId,
		param.Name,
		param.User,
		param.Email,
		param.Time,
		strings.Split(param.Codes, ",")...,
	)
	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	context.JSON(
		http.StatusOK,
		struct {
			LicenseCode string `json:"licenseCode"`
		}{
			LicenseCode: licenseCode,
		},
	)
}
