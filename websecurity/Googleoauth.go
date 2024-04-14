package websecurity

import (
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"os"
)

type Gouth struct {
	Google *oauth2.Config
}

func Init() *Gouth {
	err := godotenv.Load("application.env")
	checkerr(err)
	auth := Gouth{
		Google: &oauth2.Config{
			ClientID:     os.Getenv("google.ClientId"),
			ClientSecret: os.Getenv("google.ClientSercet"),
			Endpoint:     google.Endpoint,
			RedirectURL:  "http://localhost:8080/google/callback/",
			Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile"},
		},
	}
	return &auth
}

func checkerr(e error) {
	if e != nil {
		log.Fatal("can't open property file ")
		log.Fatal(e)
	}
}
