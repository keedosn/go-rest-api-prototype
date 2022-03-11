package handler

import (
	"net/http"

	def "git.pbiernat.dev/golang/rest-api-prototype/internal/app/definition"
	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/service"
)

var AuthLoginHandler *Handler

func init() {
	AuthLoginHandler = &Handler{
		Handle:   AuthLoginHandlerFunc,
		Request:  &def.AuthLoginRequest{},
		Response: &def.AuthLoginResponse{},
	}
}

func AuthLoginHandlerFunc(h *Handler, w http.ResponseWriter) (interface{}, int, error) {
	var req = h.Request.(*def.AuthLoginRequest)
	// u := entity.TestUser

	token, err := service.AuthService.Login(req)
	if err != nil {
		return nil, http.StatusUnauthorized, err
	}

	service.AuthService.SetCookie(w, service.AuthService.TokenCookieName, token)
	// service.AuthService.SetCookie(w, service.AuthService.RefreshTokenCookieName, refreshTtoken)

	// log.Println("user:", u, "req:", token, "err:", err)

	return nil, http.StatusOK, nil
}
