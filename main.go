package main

import (
	"fmt"
	"os"

	"bitbucket.org/joisandresky/qurban-v2/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Println("err to load local ENV", e)
	}
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		mitra := v1.Group("/mitras")
		{
			mitra.GET("/", controllers.GetMitra)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
