package main

import (
	"BACKEND_GO/internal/http/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	env := "../.env"
	err := godotenv.Load(env)
	if err != nil {
		log.Fatal("Error loading .env file on main :" + err.Error())
	}

	port := getPort()
	g := gin.New()
	routes.RegisterRoutes(g)

	log.Printf("Running on port %s", port)
	if err := g.Run(":" + port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func getPort() string {
	if port, exists := os.LookupEnv("PORT_MULTI"); exists {
		return port
	}
	if port, exists := os.LookupEnv("PORT"); exists {
		return port
	}
	return "8080"
}
