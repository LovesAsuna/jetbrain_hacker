package middleware

import "github.com/gin-gonic/gin"

func SetMiddleWare(engine *gin.Engine) {
	engine.Use(InjectCertificate)
}
