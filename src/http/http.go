package http

import "github.com/gin-gonic/gin"

func StartServer() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "OK")
	})
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
