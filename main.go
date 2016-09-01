package main

import (
	"github.com/kataras/iris"
	"testIris/controllers"
	"fmt"
)

func main() {
	routes()
	iris.Listen(":8080")
}

func middleware(c *iris.Context) {
	fmt.Println(c.Request.URI().String())
	c.Next()
}

func routes() {
	party := iris.Party("/v1", middleware)
	{
		party.API("/user", controllers.UserApi{})
	}
}

func init() {
	iris.Config.IsDevelopment = true
}