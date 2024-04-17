package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"socialnetwork/domain"
	"socialnetwork/service"
	"socialnetwork/websecurity"
)

type LoginController struct {
	auth        *websecurity.Gouth
	Serve       Controller
	userService *service.UserService
}

func (l *LoginController) Login(w http.ResponseWriter, req *http.Request) {
	fmt.Println("login")
	cookie, err := req.Cookie("sessionid")
	if err == nil {
		sessionId := cookie.Value
		l.userService.CheckSession(sessionId)
		http.Redirect(w, req, "/", http.StatusFound)
	}
	/*randomBytes := make([]byte, 32)
	_,err := rand.Read(randomBytes)
	checkerr(err)*/
	url := l.auth.Google.AuthCodeURL("")
	http.Redirect(w, req, url, http.StatusTemporaryRedirect)
}

func (l *LoginController) Callback(w http.ResponseWriter, req *http.Request) {
	code := req.FormValue("code")
	token, err := l.auth.Google.Exchange(context.TODO(), code)
	checkerr(err)
	client := l.auth.Google.Client(context.TODO(), token)
	resp, errGet := client.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	checkerr(errGet)
	defer resp.Body.Close()
	var user domain.Usr
	errJsonUser := json.NewDecoder(resp.Body).Decode(&user)
	checkerr(errJsonUser)
	l.userService.Login(&user)
	cookie := http.Cookie{
		Name:  "sessionid",
		Value: user.Id,
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, req, "/", http.StatusFound)

}

func checkerr(err error) {
	if err != nil {
		log.Fatal("err ", err)
	}
}

func (l *LoginController) Init(userService *service.UserService) {
	l.userService = userService
	l.auth = websecurity.Init()
	l.Serve.ServeMux = http.NewServeMux()
	l.Serve.HandleFunc("GET /google/login/", l.Login)
	l.Serve.HandleFunc("GET /google/callback/", l.Callback)
}
