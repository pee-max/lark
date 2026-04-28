package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lark/pkg/common/xgin"
	"lark/pkg/common/xjwt"
)

func main() {
	engine := xgin.NewGinServer()
	engine.Use(JWTAuth(), test())
	engine.Engine.GET("hello", hello)
	engine.Run(8080)
}

func hello(c *gin.Context) {
	fmt.Println("hello world!")
	c.JSON(200, "hello world!")
}

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := xjwt.ParseFromHeader(ctx)
		if err != nil {
			ctx.Abort()
			ctx.JSON(601, "parse token failed: "+err.Error())
			return
		}
		//ctx.JSON(200, "Success")
		fmt.Println("token: ", token)
		ctx.Next()
	}
}

func test() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("just a test middleware")
		ctx.Next()
		return
	}
}
