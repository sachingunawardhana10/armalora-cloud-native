package router

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sachingunawardhana10/armalora-cloud-native/internal/handler"
)

func Setup(db *pgxpool.Pool) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.Health(db))

	return mux
}
