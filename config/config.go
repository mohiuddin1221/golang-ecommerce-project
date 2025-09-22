package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations Config

type Config struct {
	Version     string
	ServiceNmae string
	HttpPort    int
	JwtToken    string
}

func loadconfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env variable", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is Required")
		os.Exit(1)
	}
	servicename := os.Getenv("SERVIC_ENAME")
	if servicename == "" {
		fmt.Println("servicename is Required")
		os.Exit(1)
	}

	jwtToken := os.Getenv("JWTTOKEN")
	if jwtToken == "" {
		fmt.Println("Jwt Token is Required")
		os.Exit(1)
	}

	http_port := os.Getenv("HTTP_PORT")
	if http_port == "" {
		fmt.Println("http_port is Required")
		os.Exit(1)
	}
	httpPort, err := strconv.Atoi(http_port)
	if err != nil {
		fmt.Println("Invalid HTTP_PORT:", err)
		os.Exit(1)
	}

	configurations = Config{
		Version:     version,
		ServiceNmae: servicename,
		HttpPort:    httpPort,
		JwtToken:    jwtToken,
	}

}

func GetConfig() Config {
	loadconfig()
	return configurations
}
