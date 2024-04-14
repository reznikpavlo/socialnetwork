package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"socialnetwork/domain"

	"log"
	"net/http"
	"socialnetwork/websecurity"
)

type LoginController struct {
	auth  *websecurity.Gouth
	Serve Controller
}

func (l *LoginController) Login(w http.ResponseWriter,
	req *http.Request) {
	fmt.Println("login")
	url := l.auth.Google.AuthCodeURL("123456789")
	http.Redirect(w, req, url, http.StatusTemporaryRedirect)
}

func (l *LoginController) Callback(w http.ResponseWriter, req *http.Request) {
	fmt.Println("callback")
	code := req.FormValue("code")
	token, err := l.auth.Google.Exchange(context.TODO(), code)
	checkerr(err)
	fmt.Println("toker is ", token)
	client := l.auth.Google.Client(context.TODO(), token)
	resp, errGet := client.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	checkerr(errGet)

	defer resp.Body.Close()
	var user domain.Usr
	fmt.Println(resp.Body)
	errJsonUser := json.NewDecoder(resp.Body).Decode(&user)
	checkerr(errJsonUser)

	fmt.Println(user)
}

func checkerr(err error) {
	if err != nil {
		log.Fatal("err")
	}
}

func (l *LoginController) Init() {
	l.auth = websecurity.Init()
	l.Serve.ServeMux = http.NewServeMux()
	l.Serve.HandleFunc("GET /google/login/", l.Login)
	l.Serve.HandleFunc("GET /google/callback/", l.Callback)
}
