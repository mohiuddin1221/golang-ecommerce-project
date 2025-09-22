package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"ecommerce/utils"
	"encoding/base64"
	"net/http"
	"strings"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// parse jwt
		// parse header and payload
		// using header and payload and secret key with hmac algorithm create signature
		// then jwt signature == newSignature

		header := r.Header.Get("Authorization")
		if header == "" {
			utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 || strings.ToLower(headerArr[0]) != "bearer" {
			utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArr[1]
		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenHeader := tokenParts[0]
		tokenPayload := tokenParts[1]
		tokenSignature := tokenParts[2]

		cnf := config.GetConfig()
		secrets := []byte(cnf.JwtToken)
		message := tokenHeader + "." + tokenPayload
		byteMessage := []byte(message)

		h := hmac.New(sha256.New, secrets)
		h.Write(byteMessage)
		hash := h.Sum(nil)
		newSignature := base64.RawURLEncoding.EncodeToString(hash)

		if tokenSignature != newSignature {
			utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
