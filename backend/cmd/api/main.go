package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/sachingunawardhana10/armalora-cloud-native/internal/config"
	"github.com/sachingunawardhana10/armalora-cloud-native/internal/database"
	"github.com/sachingunawardhana10/armalora-cloud-native/internal/router"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found")
	}

	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}
	defer db.Close()

	fmt.Println("Database connection successful")

	fmt.Println("Armalora API is running on port", cfg.AppPort)
	fmt.Println("Environment:", cfg.AppEnv)

	err = http.ListenAndServe(":"+cfg.AppPort, router.Setup(db))
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
