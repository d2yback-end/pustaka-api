package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=31072001 dbname=pustaka-api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection Error!")
	}

	fmt.Println("Postgree SQL connected successfully")

	db.AutoMigrate(&book.Book{})

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/blog", handler.BlogHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.GET("/:id", handler.DynamicHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":3000")
}
