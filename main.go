package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"warehouse/brand"
	"warehouse/config"
	"warehouse/handler"
)

var DB *gorm.DB

func main() {
	db, err := config.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	brandRepo := brand.NewRepository(db)
	brandSvc := brand.NewService(brandRepo)
	brandHandler := handler.NewBrandHandler(brandSvc)

	router := gin.Default()
	superGroup := router.Group("/api/v1")

	/* Path brand */
	brandGroup := superGroup.Group("/brand")
	brandGroup.POST("/register", brandHandler.RegisterBrandHandler)
	brandGroup.GET("/getList", brandHandler.GetBrandHandler)
	brandGroup.DELETE("/delete/:id", brandHandler.DeleteBrandHandler)

	router.Run(":5550")

}
