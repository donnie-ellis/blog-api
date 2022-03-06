package main

import (
	"github.com/gin-gonic/gin"

	"github.com/donnie-ellis/blog-api/controllers"
	"github.com/donnie-ellis/blog-api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/blogs", controllers.GetBlogs)

	r.Run()
}
