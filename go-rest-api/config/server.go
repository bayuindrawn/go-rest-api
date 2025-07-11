package config

import (
	"fmt"
	"log"
	"os"

	"go-rest-api/internal"

	"github.com/joho/godotenv"
)

func StartServer() {
	_ = godotenv.Load()

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
