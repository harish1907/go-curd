package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harish1907/curd-api/controllers"
	"github.com/harish1907/curd-api/intializers"
)


func init() {
	intializers.LocalPackageIntializers()
	intializers.ConnectionDB()
}

func main() {
	r := gin.Default()
	r.POST("/createpost", controllers.CreatePost)
	r.PUT("/postupdate/:id", controllers.PostUpdated)
	r.GET("/getallposts", controllers.GetAllPosts)
	r.GET("/getpost/:id", controllers.GetSinglePost)
	r.DELETE("/deletepost/:id", controllers.DeletePost)
	r.Run()
}