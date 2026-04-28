package xgin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GinServer struct {
	Engine *gin.Engine
}

func NewGinServer() *GinServer {
	var engine *gin.Engine
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()
	engine.Use(gin.Recovery())
	return &GinServer{Engine: engine}
}

func (g *GinServer) Run(port int) {
	addr := ":" + strconv.Itoa(port)
	err := g.Engine.Run(addr)
	if err != nil {
		fmt.Println("engine strat failed: ", err.Error())
	}
}

func (g *GinServer) Use(m ...gin.HandlerFunc) gin.IRoutes {
	return g.Engine.Use(m...)
}
