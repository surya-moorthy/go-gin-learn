package main

import (
	"go-backend-clone/restApi/grocery"
	"go-backend-clone/restApi/model"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := model.Database()

	if err != nil {
		log.Fatal(err)
	}
    db.DB()
	router := gin.Default()

	router.POST("/groceries",grocery.PostGrocery)
	router.GET("/groceries/:id",grocery.GetGrocery)
	router.PUT("/groceries/:id",grocery.UpdateGrocery)
	router.GET("/groceries",grocery.GetGroceries)
	router.DELETE("/groceries/:id",grocery.DeleteGrocery)

	router.Run(":8080")
}