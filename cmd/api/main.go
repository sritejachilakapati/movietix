package main

import (
	"encoding/json"
	"fmt"
	"log"
	"movietix/repositories"
	"movietix/services"
)

func main() {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)

	user, err := userService.GetUserByID(1)

	if err != nil {
		log.Fatalf("Error")
	}

	jsonData, _ := json.Marshal(user)

	jsonUser := string(jsonData)
	fmt.Println(jsonUser)
}
