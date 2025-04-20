package middleware

import (
	"fmt"
	"github.com/LovesAsuna/jetbrains_hacker/internal/cert"
	"github.com/LovesAsuna/jetbrains_hacker/server/config"
	"github.com/LovesAsuna/jetbrains_hacker/server/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sync"
)

var certCache = struct {
	userCert          *cert.Certificate
	licenseServerCert *cert.Certificate
	rwLock            sync.RWMutex
}{}

func InjectCertificate(context *gin.Context) {
	certCache.rwLock.RLock()
	if certCache.userCert == nil || certCache.licenseServerCert == nil {
		certCache.rwLock.RUnlock()
		certCache.rwLock.Lock()
		if certCache.userCert != nil && certCache.licenseServerCert != nil {
			certCache.rwLock.Unlock()
		} else {
			var err error
			isExist := func(filePath string) bool {
				_, err := os.Stat(filePath)
				return err == nil || !os.IsNotExist(err)
			}
			if isExist(config.Config.UserCertPath) && isExist(config.Config.UserPrivateKeyPath) {
				certCache.userCert, err = cert.CreateCertFromFile(config.Config.UserCertPath, config.Config.UserPrivateKeyPath)
				if err != nil {
					certCache.rwLock.Lock()
					_ = context.AbortWithError(http.StatusInternalServerError, err)
					return
				}
			} else {
				if certCache.userCert, err = cert.GenerateFakeCertificate(
					cert.JetProfileCert.CommonName(),
					"create by license server",
					config.Config.UserCertPath,
					config.Config.UserPrivateKeyPath,
				); err != nil {
					certCache.rwLock.Lock()
					_ = context.AbortWithError(http.StatusInternalServerError, err)
					return
				}
			}

			if isExist(config.Config.LicenseServerCertPath) && isExist(config.Config.LicenseServerPrivateKeyPath) {
				certCache.licenseServerCert, err = cert.CreateCertFromFile(config.Config.LicenseServerCertPath, config.Config.LicenseServerPrivateKeyPath)
				if err != nil {
					certCache.rwLock.Lock()
					_ = context.AbortWithError(http.StatusInternalServerError, err)
					return
				}
			} else {
				if certCache.licenseServerCert, err = cert.GenerateFakeCertificate(
					cert.LicenseServerCert.CommonName(),
					fmt.Sprintf("%s.lsrv.jetbrains.com", "license_server"),
					config.Config.LicenseServerCertPath,
					config.Config.LicenseServerPrivateKeyPath,
				); err != nil {
					certCache.rwLock.Lock()
					_ = context.AbortWithError(http.StatusInternalServerError, err)
					return
				}
			}
			certCache.rwLock.Unlock()
		}
	} else {
		certCache.rwLock.RUnlock()
	}

	certPool := &handler.CertPool{
		UserCert:          certCache.userCert,
		LicenseServerCert: certCache.licenseServerCert,
	}
	context.Set(handler.CertPoolKey, certPool)
}
