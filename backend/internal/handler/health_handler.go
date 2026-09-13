package handler

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sachingunawardhana10/armalora-cloud-native/internal/database"
)

func Health(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := database.HealthCheck(db)

		if err != nil {
			http.Error(w, "Database is unavailable", http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Armalora API is healthy"))
	}
}
