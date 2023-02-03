package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"warehouse/brand"
	"warehouse/handler"
)

var DB *gorm.DB

func main() {
	var err error
	dsn := "host=localhost user=postgres password=Berandallasta15 dbname=db_warehouse port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = db

	if err != nil {
		fmt.Println("success to connection database")
	}

	brandRepository := brand.NewRepository(db)
	brandService := brand.NewService(brandRepository)

	brandHandler := handler.NewBrandHandler(brandService)

	router := gin.Default()
	api := router.Group("/api/v3")
	api.POST("/registerbrand", brandHandler.RegisterBrandHandler)
	api.GET("/getbrandlist", brandHandler.GetBrandHandler)
	api.DELETE("/deletebrand", brandHandler.DeleteBrandHandler)
	router.Run(":5550")

}
