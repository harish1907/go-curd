package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/harish1907/curd-api/intializers"
	"github.com/harish1907/curd-api/models"
)

func CreatePost(c *gin.Context) {
	// Get data in Body
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	c.Bind(&body)

	// Create a Post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := intializers.DB.Create(&post)

	// return a response
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"result": post,
	})
}

func GetAllPosts(c *gin.Context) {
	// Get all the post using gorm
	var posts []models.Post
	intializers.DB.Find(&posts)

	//return json response....
	c.JSON(200, gin.H{
		"result": posts,
	})
}

func GetSinglePost(c *gin.Context) {
	// Get value in URLS..
	id := c.Param("id")

	// Search By id..
	var post models.Post
	intializers.DB.Find(&post, id)

	// Return Response..
	c.JSON(200, gin.H{
		"result": post,
	})
}

func PostUpdated(c *gin.Context) {
	// Get the id of the URL
	id := c.Param("id")

	// Get the data of request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the post were updated
	var post models.Post
	intializers.DB.Find(&post, id)

	// update it
	intializers.DB.Model(&post).Updates(
		models.Post{
			Title: body.Title,
			Body:  body.Body,
		})

	// Return response
	c.JSON(200, gin.H{
		"message": "Update successfully",
		"Post":    post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	intializers.DB.Delete(&models.Post{}, id)
	c.Status(200)
}
