package server

import (
	"github.com/LovesAsuna/jetbrains_hacker/server/config"
	"github.com/LovesAsuna/jetbrains_hacker/server/middleware"
	"github.com/LovesAsuna/jetbrains_hacker/server/router"
	"github.com/gin-gonic/gin"
)

func RunServer() error {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	middleware.SetMiddleWare(engine)
	router.SetRouter(engine)
	return engine.Run(config.Config.Addr)
}
