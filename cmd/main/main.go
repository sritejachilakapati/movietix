package main

import (
	"context"
	"fmt"
	"log"
	"movietix/internal/config"
	"movietix/internal/database"
	"movietix/models"
	"movietix/repositories"
	"movietix/services"

	"github.com/google/uuid"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ctx := context.Background()

	db, err := database.InitDB(ctx)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	user := &models.User{
		ID:       uuid.New().String(),
		Name:     "John Doe",
		Email:    "foo@bar.com",
		Password: "foobar123",
		IsActive: true,
		Role:     "admin",
	}

	err = userService.CreateUser(ctx, user)

	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	fmt.Println("User created successfully")

	fetchedUser, err := userService.GetUserByID(ctx, user.ID)

	if err != nil {
		log.Fatalf("Error fetching user: %v", err)
	}

	fmt.Printf("Fetched user: %+v\n", fetchedUser)
}
