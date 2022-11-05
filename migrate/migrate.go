package main

import (
	"github.com/bimaagung/go-crud/initializers"
	"github.com/bimaagung/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
