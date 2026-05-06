package server

import (
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/router"
	"lark/pkg/common/xgin"
)

type Server struct {
	cfg       *config.Config
	ginServer *xgin.GinServer
}

func NewServer() *Server {
	s := &Server{cfg: config.GetConfig(), ginServer: xgin.NewGinServer()}
	return s
}

func (s *Server) Run() {
	router.Register(s.ginServer.Engine)
	s.ginServer.Run(s.cfg.Port)
}
