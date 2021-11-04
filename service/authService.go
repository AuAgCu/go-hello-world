package service

import (
	firebaseUtil "app/firebase"
	"context"
	"fmt"
	"log"
)

var AuthServiceSingleton = &authServiceImpl{}

type AuthService interface {
	Verify(jwtToken string) bool
}

type authServiceImpl struct {
}

func (authservice authServiceImpl) Verify(jwtToken string) bool {
	println(jwtToken)
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
