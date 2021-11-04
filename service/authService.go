package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

var AuthServiceSingleton = &authServiceImpl{}

type AuthService interface {
	Verify(jwtToken string) bool
	GetUserId(jwtToken string) string
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

	token, err := auth.VerifyIDToken(context.Background(), jwtToken)
	if err != nil {
		fmt.Printf("error verifying ID token: %v\n", err)
		return false
	}

	userId := authservice.GetUserId(jwtToken)

	log.Printf("Verified ID token: %v\n", token)
	log.Printf("Verified ID token: %s\n", userId)
	return true
}

func (authService authServiceImpl) GetUserId(jwtToken string) string {
	hcs, err := authService.decomposeFB(jwtToken)
	if err != nil {
		log.Fatalln("Error : ", err)
	}

	payload, err := authService.decodeClaimFB(hcs[1])
	if err != nil {
		log.Fatalln("Error : ", err)
	}

	return payload.Subject
}

//DecomposeFB  JWTをHeader, claims, 署名に分解
func (_ authServiceImpl) decomposeFB(jwt string) ([]string, error) {
	hCS := strings.Split(jwt, ".")
	if len(hCS) == 3 {
		return hCS, nil
	}
	return nil, errors.New("Error jwt str decompose: inccorrect number of segments")

}

type FireBaseCustomToken struct {
	auth.Token
	Email string `json:"email"`
}

//DecodeClaimFB JWTのclaims部分をFireBaseCustomTokenの構造体にデコード
func (_ authServiceImpl) decodeClaimFB(payload string) (*FireBaseCustomToken, error) {
	payloadByte, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return nil, errors.New("Error jwt token decode: " + err.Error())
	}

	var tokenJSON FireBaseCustomToken
	err = json.Unmarshal(payloadByte, &tokenJSON)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, errors.New("Error jwt token unmarshal: " + err.Error())
	}

	return &tokenJSON, nil
}
