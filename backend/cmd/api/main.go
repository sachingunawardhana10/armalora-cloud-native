package main

import (
	"fmt"
	"net/http"

	"github.com/sachingunawardhana10/armalora-cloud-native/internal/config"
)

func main() {
	cfg := config.Load()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Armalora API")
	})

	fmt.Println("Armalora API is running on port", cfg.AppPort)
	fmt.Println("Environment:", cfg.AppEnv)

	err := http.ListenAndServe(":"+cfg.AppPort, nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
