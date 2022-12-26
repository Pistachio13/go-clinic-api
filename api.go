package main

import (
	"clinic/go-clinic-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/go_clinic?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	r := gin.Default()
	r.GET("/patients", func(c *gin.Context) {
    var patients []model.Patient
    db.Find(&patients)
		c.JSON(http.StatusOK, patients)
	})

 	r.GET("/patients/:id", func(c *gin.Context) {
    id := c.Param("id")
    var patient model.Patient 
    db.First(&patient, id)
    c.JSON(http.StatusOK, patient)
  })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
