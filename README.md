# Rotator

[![Go Reference](https://pkg.go.dev/badge/github.com/dimasadyaksa/rotator.svg)](https://pkg.go.dev/github.com/dimasadyaksa/rotator)
[![Go Report Card](https://goreportcard.com/badge/github.com/dimasadyaksa/rotator)](https://goreportcard.com/report/github.com/dimasadyaksa/rotator)
[![License](https://img.shields.io/github/license/dimasadyaksa/rotator)](LICENSE)

## Overview
Rotator is a Go package that provides a fundamental implementation of a rotating container designed to hold elements of various types. 
This package facilitates efficient management and manipulation of elements through key-based operations, all while ensuring Go routine safety. 
It is particularly well-suited for use cases where cyclic traversal of element lists or rapid access by their respective keys is necessary.

## Installation

To install the package, simply run:

```shell
go get github.com/dimasadyaksa/rotator
```

## Usage



Here is an example of using Rotator as a key rotator for JWT signing.

First, make sure to install the "github.com/golang-jwt/jwt" package if you haven't already:

```shell
go get github.com/golang-jwt/jwt
```

```go
package main

import (
	"fmt"
	"github.com/dimasadyaksa/rotator"
	"github.com/golang-jwt/jwt"
)

// JWTKey Custom struct to represent JWT signing keys.
type JWTKey struct {
	ID  string
	Key string
}

func main() {
	// Create a Rotator for JWT signing keys.
	keyRotator := rotator.New(func(key JWTKey) string {
		return key.ID
	})

	keys := []JWTKey{
		{
			ID:  "key1",
			Key: "key1_secret",
		},
		{
			ID:  "key2",
			Key: "key2_secret",
		},
	}

	// Add JWT signing keys to the rotator.
	keyRotator.Add(keys...)

	// Get the current key from the rotator.
	currentKey, err := keyRotator.Rotate()
	if err != nil {
		fmt.Println("Error getting current key:", err)
		return
	}

	// Create a new JWT token.
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the "kid" (Key ID) in the JWT header.
	token.Header["kid"] = currentKey.ID

	// Sign the token with the current key.
	tokenString, err := token.SignedString([]byte(currentKey.Key))
	if err != nil {
		fmt.Println("Error signing token:", err)
		return
	}

	// Print the JWT token.
	fmt.Println("JWT Token:")
	fmt.Println(tokenString)
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/dimasadyaksa/rotator/blob/master/LICENSE) file for details.
