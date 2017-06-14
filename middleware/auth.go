package middleware

import (
	"net/http"

	"github.com/greg-gb/go-rest-api-starter/utils"
)

var (
	authorizationHeader = "Authorization"
)

type middleware func(http.Handler) http.Handler

func CheckAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for name, headers := range r.Header {
			if name == authorizationHeader {
				for _, header := range headers {
					// Check the Authorization Token. For testing purposes "test" is our token.
					// github.com/dgrijalva/jwt-go
					if header != "test" {
						utils.SendRes(w, r, &utils.Res{401, "Authorization token is invalid. Try test :)"})
						return
					}

					h.ServeHTTP(w, r)
					return
				}
			}
		}
		utils.SendRes(w, r, &utils.Res{400, "Authorization header is required."})
		return
	})
}
