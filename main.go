package main

import (
	"github.com/RohitDeepati/go-ecomerce-BE/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.HealthRoute(router)
	router.Run(":8080")
}
