package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"practicing-amazon-s3-golang/controllers"
	"time"
)

func main() {

	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	}))

	router.POST("/upload",controllers.Upload)
	router.Run()

}