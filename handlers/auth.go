package handlers

import (
	"net/http"

	"github.com/greg-gb/go-rest-api-starter/utils"
)

func Auth() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			utils.SendRes(w, r, &utils.Res{200, "OK"})
			return
		}
		utils.SendRes(w, r, &utils.Res{405, "Method not allowed."})
		return
	})
}
