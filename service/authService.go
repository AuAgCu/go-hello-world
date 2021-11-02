package service

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
)

type AuthService interface {
	Verify(jwtToken string) bool
}

type AuthServiceImpl struct {
}

func NewAuthService() AuthService {
	return &AuthServiceImpl{}
}

func (authservice AuthServiceImpl) Verify(jwtToken string) bool {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		return false
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return false
	}

	token, err := auth.VerifyIDToken(context.Background(), jwtToken)
	if err != nil {
		fmt.Printf("error verifying ID token: %v\n", err)
		return false
	}

	log.Printf("Verified ID token: %v\n", token)
	return true
}
