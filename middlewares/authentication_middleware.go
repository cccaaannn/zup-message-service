package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"zup-message-service/data/dtos"
	"zup-message-service/data/enums"
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
			token := split[1]
			tokenPayloadResult := services.IsAuthorized(token)
			if !tokenPayloadResult.Status {
				message := dtos.DataResult[dtos.TokenPayload]{Status: false, Message: "Forbidden", Data: nil}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(message)
			} else {

				// Add token and payload to context to access them on services
				ctx := context.WithValue(r.Context(), enums.TOKEN, token)
				ctx = context.WithValue(ctx, enums.TOKEN_PAYLOAD, tokenPayloadResult.Data)

				next.ServeHTTP(w, r.WithContext(ctx))
			}
		}
	})
}
