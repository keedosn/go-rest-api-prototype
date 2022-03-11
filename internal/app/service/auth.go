package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/config"
	def "git.pbiernat.dev/golang/rest-api-prototype/internal/app/definition"
	"github.com/golang-jwt/jwt"
)

var (
	AuthService *Auth

	ErrUserNotFound = errors.New("user not found")
	ErrTokenError   = errors.New("failed to generate JWT token")
)

func init() {
	expire, _ := strconv.Atoi(config.GetEnv("AUTH_TOKEN_EXPIRE_TIME", "5"))
	secret := []byte(config.GetEnv("AUTH_SECRET_HMAC", "B413IlIv9nKQfsMCXTE0Cteo4yHgUEfqaLfjg73sNlh"))

	AuthService = &Auth{expire, "jwt_token", "jwt_token_refresh", secret}
}

type Auth struct {
	ExpireTime             int // token expire time in minutes
	TokenCookieName        string
	RefreshTokenCookieName string

	secret []byte // signing key
}

func (a *Auth) Login(r *def.AuthLoginRequest) (string, error) {
	if r.Username == "admin" && r.Password == "secret" {
		token, err := a.createToken()
		if err != nil {
			return "", ErrTokenError
		}

		return token, nil
	}

	return "", ErrUserNotFound
}

// SetCookie appends cookie header to response
func (a *Auth) SetCookie(w http.ResponseWriter, name, token string) {
	c := &http.Cookie{
		Name:   name,
		Value:  token,
		MaxAge: a.ExpireTime * 60,
		Path:   "/",
	}
	http.SetCookie(w, c)
}

func (a Auth) createToken() (string, error) {
	// log.Println("now:", time.Now().Unix())
	// log.Println("expire at:", time.Now().Add(time.Duration(a.ExpireTime)*time.Minute).Unix())
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(a.ExpireTime) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.secret)
}

func (a *Auth) validateToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return a.secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println(claims)
	} else {
		return err
	}

	return nil
}

func (a Auth) ValidateUserTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cToken, err := r.Cookie(a.TokenCookieName)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(def.Error("Missing JWT Token cookie"))

			return
		}

		if err := a.validateToken(cToken.Value); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(def.Error(err.Error()))

			return
		}

		next.ServeHTTP(w, r)
	})
}
