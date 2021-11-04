package service

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

var AuthServiceSingleton = &authServiceImpl{}

type AuthService interface {
	Verify(jwtToken string) bool
}

type authServiceImpl struct {
}

func (authservice authServiceImpl) Verify(jwtToken string) bool {
	credentials, err := google.CredentialsFromJSON(context.Background(), []byte(os.Getenv("FIREBASE_KEYFILE_JSON")))
	if err != nil {
		log.Printf("error credentials from json: %v\n", err)
	}
	opt := option.WithCredentials(credentials)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error initializing app: %v\n", err)
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return false
	}

	println("huga")
	token, err := auth.VerifyIDToken(context.Background(), jwtToken)
	if err != nil {
		fmt.Printf("error verifying ID token: %v\n", err)
		return false
	}

	log.Printf("Verified ID token: %v\n", token)
	return true
}
