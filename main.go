package main

import (
	"ecommerce/utils"
	"fmt"
)

// "ecommerce/cmd"

func main() {
	// text := "topmohiudinn"
	// bytecode := []byte(text)
	// fmt.Println("Bytecode", bytecode)
	// encoded := base64.RawURLEncoding.EncodeToString(bytecode)
	// fmt.Println("URL Encoding without padding:", encoded)

	// text2 := []byte("hello")
	// hash := sha256.Sum256(text2)
	// fmt.Println("URL Encoding without padding:", hash)
	// secrets := []byte("my_secret")
	// message := []byte("helloworld")

	// h := hmac.New(sha256.New, secrets)
	// h.Write(message)
	// signature := h.Sum(nil)
	// fmt.Println("signature..........", signature)

	jwt, err := utils.CreateJwt(utils.JWTPayload{
		Sub:         100,
		FirstName:   "Mohiuddin",
		LastName:    "Topu",
		Email:       "topu@gmail.com",
		IsShopOwner: false,
	}, "my_secret")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jwt)

	// cmd.Serve()
}
