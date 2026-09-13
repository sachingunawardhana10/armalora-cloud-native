package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/sachingunawardhana10/armalora-cloud-native/internal/config"
	"github.com/sachingunawardhana10/armalora-cloud-native/internal/database"
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
	defer db.Close(context.Background())

	fmt.Println("Database connection successful")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Armalora API")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		err := database.HealthCheck(db)

		if err != nil {
			http.Error(w, "Database is unavailable", http.StatusServiceUnavailable)
			return
		}

		fmt.Fprintln(w, "Armalora API is healthy")
	})

	fmt.Println("Armalora API is running on port", cfg.AppPort)
	fmt.Println("Environment:", cfg.AppEnv)

	err = http.ListenAndServe(":"+cfg.AppPort, nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
