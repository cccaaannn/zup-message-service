package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"
	"zup-message-service/dtos"
	"zup-message-service/services"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawToken := r.Header.Get("authorization")
		split := strings.Split(rawToken, " ")

		if len(split) != 2 {
			message := dtos.DataResult[dtos.TokenPayload]{Status: false, Message: "Forbidden", Data: nil}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(message)
		} else {
			user := services.IsAuthorized(split[1])
			if !user.Status {
				message := dtos.DataResult[dtos.TokenPayload]{Status: false, Message: "Forbidden", Data: nil}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(message)
			} else {
				next.ServeHTTP(w, r)
			}
		}
	})
}
