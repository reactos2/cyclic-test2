package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	port := os.Getenv("PORT")
	fmt.Println("port check 1: ", port)
	if port == "" {
		port = "8080"
	}
	fmt.Println("port check 2: ", port)

	r := gin.Default()

	r.GET("/", handleHomePage)
	r.GET("/greeter/greet", handleGreeter)
	r.GET("/test2", handleTest2)

	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello, %s!", name),
		})
	})

	fmt.Println("port check 3: ", port)
	r.Run(fmt.Sprintf(":%s", port))
}

func handleHomePage(ctx *gin.Context) {
	ctx.String(200, "welcome to koyeb test 1")
}
func handleGreeter(ctx *gin.Context) {
	ctx.String(200, "welcome to koyeb test 1")
}
func handleTest2(ctx *gin.Context) {
	fmt.Println("http get test2")
	ctx.String(200, "koyeb test 2")
}
