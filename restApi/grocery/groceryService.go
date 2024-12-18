package grocery

import (
	"go-backend-clone/restApi/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewGrocery struct {
	Name     string `json:"name" binding:required`
	Quantity string `jaon:"quantity" binding:required`
}

type Updategrovery struct {
	Name     string `json:"name" binding:required`
	Quantity string `json:"quantity" binding:required`
}

func GetGroceries(c *gin.Context) {
	var groceries []model.Grocery

	db, err := model.Database()

	if err != nil {
		log.Fatal(err)
	}
	if err := db.Find(&groceries).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, groceries)
}

func PostGroceries(c *gin.Context){
	 var grocery NewGrocery

	 if err := c.ShouldBindJSON(&grocery); err != nil {
		log.Fatal(err)
	 }

	 newGrocery := model.Grocery{Name: grocery.Name , Quantity: grocery.Quantity}

	 db , err := model.Database()
	 if err != nil {
		log.Fatal(err)
	 }
	
	 if err := db.Create(&newGrocery).Error ; err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error" : err.Error()})
		return 
	 }
	 c.JSON(http.StatusOK , newGrocery)
}