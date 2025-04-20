package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Fs(context *gin.Context) {
	context.FileFromFS(context.Request.URL.Path, http.FS(os.DirFS("static")))
}
