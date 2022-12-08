package main

import (
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/blog", handler.BlogHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.GET("/:id", handler.DynamicHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":3000")
}
