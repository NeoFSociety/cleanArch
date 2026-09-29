package main

import (
	"context"
	"log"
	"net/http"

	"github.com/NeoFSociety/cleanArch/internal/repository/postgres"
	"github.com/NeoFSociety/cleanArch/internal/repository/postgres/pool"
	"github.com/NeoFSociety/cleanArch/internal/service"
	"github.com/NeoFSociety/cleanArch/internal/transport"
)

func main() {
	cfg := pool.NewConfigMust()

	ctx := context.Background()

	dbPool, err := pool.NewPool(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to create pool: %v", err)
	}
	defer dbPool.Close()

	userRepo := postgres.NewUserRepo(dbPool)
	userService := service.NewUserService(userRepo)
	userHandler := transport.NewUserHandler(userService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", userHandler.Register)
	mux.HandleFunc("GET /users/{uuid}", userHandler.GetUserByID)
	mux.HandleFunc("DELETE /users/{uuid}", userHandler.Delete)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
