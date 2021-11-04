package firebaseUtil

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

var FirebaseUtil = &firebaseUtil{}

type firebaseUtil struct {
	App firebase.App
}

func (firebaseUtil firebaseUtil) InitFirebase() {
	credentials, err := google.CredentialsFromJSON(context.Background(), []byte(os.Getenv("FIREBASE_KEYFILE_JSON")))
	if err != nil {
		log.Printf("error credentials from json: %v\n", err)
	}
	opt := option.WithCredentials(credentials)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error initializing app: %v\n", err)
	}

	println(app)

	firebaseUtil.App = *app
	println("firebase init")
}
