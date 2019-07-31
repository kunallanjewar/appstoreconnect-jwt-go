# AppStoreConnect-JWT-Go

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![godoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?)](https://godoc.org/github.com/kunallanjewar/appstoreconnect-jwt-go#Client.BearerToken)
[![Build Status](https://travis-ci.org/kunallanjewar/appstoreconnect-jwt-go.svg?branch=master)](https://travis-ci.org/kunallanjewar/appstoreconnect-jwt-go)
[![Coverage Status](https://coveralls.io/repos/github/kunallanjewar/appstoreconnect-jwt-go/badge.svg)](https://coveralls.io/github/kunallanjewar/appstoreconnect-jwt-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/kunallanjewar/appstoreconnect-jwt-go)](https://goreportcard.com/report/github.com/kunallanjewar/appstoreconnect-jwt-go)

AppStoreConnect-JWT-Go is a Go package that provides an easy way generate JWT token needed for accessing AppStoreConnect API.

ApStoreConnect JWT Token [requirements](https://developer.apple.com/documentation/appstoreconnectapi/generating_tokens_for_api_requests).

## Example

[main.go](example/main.go)

```golang
        client, err := jwt.New(cfg)
	if err != nil {
		panic(err)
	}

	tokenString, err := client.BearerToken()
	if err != nil {
		panic(err)
	}

	fmt.Println(tokenString)
```

Result:

```
$ go run example/main.go
eyJhbGciOiJFUzI1NiIsImtpZCI6IjJYOVI0SFhGMzQiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJhcHBzdG9yZWNvbm5lY3QtdjEiLCJleHAiOjE1NjQ1OTQ0NzEsImlhdCI6MTU2NDU5Mzg3MSwiaXNzIjoiNTcyNDY1NDItOTZmZS0xYTYzLWUwNTMtMDgyNGQwMTEwNzJhIn0.Tpqv1ZoDcv7CsDaq4ZF8bycN3hJexYrBQbzsUEd6hNV94bQ_gIES1nsCDlF9-JMrWlT7sa
1ET2kZcWBezfUe5w
```

## Todo

- [x] A few more tests
- [ ] Remove dependency on [jwt-go](https://github.com/dgrijalva/jwt-go) package
- [ ] Document more code
