package router

import (
	"github.com/LovesAsuna/jetbrains_hacker/server/handler"
	"github.com/gin-gonic/gin"
)

func SetRouter(engine *gin.Engine) {
	engine.Handle("GET", "/rpc/obtainTicket.action", handler.ObtainTicket)
	engine.Handle("GET", "/rpc/ping.action", handler.Ping)
	engine.Handle("GET", "/rpc/releaseTicket.action", handler.ReleaseTicket)
	engine.Handle("GET", "/config/:type", handler.Config)
	engine.LoadHTMLGlob("static/*")
	engine.NoRoute(handler.Fs)
}
