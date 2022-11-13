package main

import (
	"github.com/harish1907/curd-api/intializers"
	"github.com/harish1907/curd-api/models"
)

func init() {
	intializers.LocalPackageIntializers()
	intializers.ConnectionDB()
}

func main() {
	intializers.DB.AutoMigrate(&models.Post{})
}
