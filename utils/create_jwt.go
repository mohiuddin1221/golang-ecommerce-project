package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type JWTHeader struct {
	Alg string `json:"alg"` // Algorithm
	Typ string `json:"typ"` // Token type
}

type JWTPayload struct {
	Sub         int    `json:"sub"`
	FirstName   string `json:"first_name"`    // First Name
	LastName    string `json:"last_name"`     // Last Name
	Email       string `json:"email"`         // Email address
	IsShopOwner bool   `json:"is_shop_owner"` // Optional: Shop owner flag
}

func CreateJwt(payload JWTPayload, secret string) (string, error) {
	header := JWTHeader{
		Alg: "HS256",
		Typ: "JWT",
	}
	headerJson, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerBase64 := base64.RawURLEncoding.EncodeToString(headerJson)

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	payloadbase64 := base64.RawURLEncoding.EncodeToString(payloadJson)

	// create signature
	message := headerBase64 + "." + payloadbase64
	bytecode := []byte(secret)
	h := hmac.New(sha256.New, bytecode)
	h.Write([]byte(message))
	signature := h.Sum(nil)
	signaturebase64 := base64.RawURLEncoding.EncodeToString(signature)
	jwt := message + "." + signaturebase64

	return jwt, nil

}
