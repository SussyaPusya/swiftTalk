package transport

import (
	"github.com/gin-gonic/gin"
)

type router struct {
	// s Service
	m *Middleware
}

func NewRouter(m *Middleware) *router {
	return &router{m: m}
}

func (r *router) Run() {
	ginRouter := gin.Default()

	ginRouter.Use(r.m.AuthMiddleware())
	ginRouter.Run(":8082")
}
