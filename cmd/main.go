package main

import (
	"context"
	"log"

	"github.com/sritejachilakapati/movietix/internal/config"
	"github.com/sritejachilakapati/movietix/internal/database"
	"github.com/sritejachilakapati/movietix/internal/repository"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	ctx := context.Background()
	conn := database.Connect(ctx)
	defer conn.Close(ctx)

	userRepo := repository.New(conn)

	users, err := userRepo.GetAllUsers(ctx)
	if err != nil {
		log.Fatalf("Error fetching users: %v", err)
	}

	for _, user := range users {
		log.Println(user)
	}

}
