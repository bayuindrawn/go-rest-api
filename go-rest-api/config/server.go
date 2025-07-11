package config

import (
	"fmt"
	"log"
	"os"

	"go-rest-api/internal"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func StartServer() {
	_ = godotenv.Load()
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode // fallback
	}
	gin.SetMode(ginMode)

	db, err := ConnectDB()
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	defer db.Close()

	app := internal.Init(db)
	router := SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running at http://localhost:" + port)
	log.Fatal(router.Run(":" + port))
}
