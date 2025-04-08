package main

import (
	"context"
	"log"

	"github.com/dinizgab/split-api/internal/database"
	"github.com/dinizgab/split-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.New()
	if err != nil {
		log.Println(err)
	}
	defer db.Close(context.Background())

	router := gin.Default()
	router.POST("/user/register", handlers.RegisterUser(db))
	router.GET("/group/:id", handlers.GetGroupByID(db))
	router.POST("/group", handlers.CreateNewGroup(db))
	router.POST("/group/:id/user", handlers.AddUserToGroup(db))

	log.Println(router.Run(":8000"))
}
