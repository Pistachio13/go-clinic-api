package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  "clinic/go-clinic-api/model"
)

func main() {
  dsn := "root@tcp(127.0.0.1:3306)/go_clinic?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(&model.Patient{})

  // Create
  db.Create(&model.Patient{Fname: "John", Lname: "Doe", Tel: "0568889423", Diagnosis: "Stroke"})
  db.Create(&model.Patient{Fname: "Jane", Lname: "Doe", Tel: "0564814495", Diagnosis: "HNP"})


}