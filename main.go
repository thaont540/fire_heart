package main

import (
	"fire_heart/controllers"
	"fire_heart/models/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connection()
	userController := new (controllers.UserController)
	router := gin.Default()
	router.POST("/", userController.Store)
	router.GET("/:email", userController.Show)
	router.DELETE("/:email", userController.Delete)
	//client := db.Client.Database("fire_heart").Collection("posts")
	//filter := bson.D
	//type Post struct {
	//
	//	Title string `json:"title,omitempty""`
	//
	//	Body string `json:"body,omitempty""`
	//
	//}
	//
	//var post Post

	//fmt.Println(*client)

	_ = router.Run()
}
