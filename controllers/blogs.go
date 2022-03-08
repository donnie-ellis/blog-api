// controlers/books.go

package controllers

import (
	"net/http"
	"time"

	"github.com/donnie-ellis/blog-api/models"
	"github.com/gin-gonic/gin"
)

type CreateBlogInput struct {
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
}

type ModifyBlogInput struct {
	Title    string    `json:"title"`
	Text     string    `json:"text"`
	Modified time.Time `json:modified`
}

// GET /blogs
// Get all blogs
func GetBlogs(c *gin.Context) {
	var blogs []models.Blog
	models.DB.Find(&blogs)

	c.JSON(http.StatusOK, gin.H{"data": blogs})
}

// POST /blogs
// Create a new blog
func CreateBlog(c *gin.Context) {
	var input CreateBlogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()

	blog := models.Blog{Title: input.Title, Text: input.Text, Created: now, Modified: now}
	models.DB.Create(&blog)

	c.JSON(http.StatusOK, gin.H{"data": blog})
}

//PATCH /blogs/:id
func ModifyBlog(c *gin.Context) {
	var blog models.Blog
	if err := models.DB.Where("id = ?", c.Param("id")).First(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Blog not found"})
		return
	}

	var input ModifyBlogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&blog).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": blog})
}

//DELETE /blogs/:id
// Delete a blog
func DeleteBlog(c *gin.Context) {
	var blog models.Blog
	if err := models.DB.Where("id = ?", c.Param("id")).First(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Blog not found"})
		return
	}

	models.DB.Delete(&blog)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
