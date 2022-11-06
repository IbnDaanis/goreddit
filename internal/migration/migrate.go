package main

import (
	"github.com/ibndaanis/golang-rest/internal/initializers"
	"github.com/ibndaanis/golang-rest/internal/models"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDataBase()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
