package firebase

import (
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"context"
	"log"

	"fmt"
)

var FirebaseApp *firebase.App

func InitializeAppDefault(){
	// [START initialize_app_default_golang]
	opt := option.WithCredentialsFile("../service-account-file.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	} else {
		fmt.Println("initialize success")
	}
	// [END initialize_app_default_golang]

	FirebaseApp = app
}

