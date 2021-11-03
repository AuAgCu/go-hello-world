package service

import (
	firebaseUtil "app/firebase"
	"context"
	"fmt"
	"log"
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

	auth, err := firebaseUtil.FirebaseUtil.App.Auth(context.Background())
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
