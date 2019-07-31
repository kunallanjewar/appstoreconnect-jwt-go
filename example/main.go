package main

import (
	"fmt"
	"time"

	jwt "github.com/kunallanjewar/appstoreconnect-jwt-go"
)

// PrivateKey is the key provided by AppStoreConnect.
const PrivateKey = `
-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgLr73kREgj9lV5HHg
dFDBpMfzJ/Y/hdielkVexW5ML9OhRANCAASxG1wW2Wlth3tE3fxjAAf0bd83M1p8
m4B8ipZ3jY5hvTb9zbM8GbhExZotyZW/B27acARhMToQcFIpO3GurIrd
-----END PRIVATE KEY-----
`

func main() {
	cfg := &jwt.Config{
		ISS:       "57246542-96fe-1a63-e053-0824d011072a",
		KID:       "2X9R4HXF34",
		ExpiresIn: 10 * time.Minute,
		AUD:       "appstoreconnect-v1",
		PK:        PrivateKey,
	}

	client, err := jwt.New(cfg)
	if err != nil {
		panic(err)
	}

	tokenString, err := client.BearerToken()
	if err != nil {
		panic(err)
	}

	fmt.Println(tokenString)
}
