package main

import (
	auth "api_gateway/usecase"
	"fmt"
)

func main() {
	
	login := auth.NewLogin()

	username := "admin"
	password := "admin123"

	isAuthenticated := login.Authentication(username, password)

	if isAuthenticated {
		fmt.Println("Login successful")
	} else {
		fmt.Println("Invalid username or password")
	}
}

