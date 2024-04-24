package controller

import (
	"log"
	"net/http"
	"socialnetwork/service"
	"socialnetwork/websecurity"
)

type LogoutController struct {
	auth        *websecurity.Gouth
	Serve       Controller
	userService *service.UserService
}

func (l *LogoutController) Logout(w http.ResponseWriter, req *http.Request) {
	log.Println("/logout")
	cookie, err := req.Cookie("sessionid")
	if err == nil {
		sessionId := cookie.Value
		l.userService.DeleteSession(sessionId)
		cookie.Value = ""
		cookie.MaxAge = 0
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, req, "/", http.StatusFound)
}

func (l *LogoutController) Init(userService *service.UserService) {
	l.userService = userService
	l.auth = websecurity.Init()
	l.Serve.ServeMux = http.NewServeMux()
	l.Serve.HandleFunc("GET /logout", l.Logout)

}
