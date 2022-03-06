// controlers/books.go

package controllers

import (
	"net/http"

	"github.com/donnie-ellis/blog-api/models"
	"github.com/gin-gonic/gin"
)

// GET /blogs
// Get all blogs
func GetBlogs(c *gin.Context) {
	var blogs []models.Blog
	models.DB.Find(&blogs)

	c.JSON(http.StatusOK, gin.H{"data": blogs})
}
