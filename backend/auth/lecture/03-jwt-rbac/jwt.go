package main

import (
	//...
	// import the jwt-go library
	"github.com/dgrijalva/jwt-go"
	//...
)

// Jwt key yang akan dipakai untuk membuat signature
var jwtKey = []byte("key")

// Data user - password yang bisa mengakses api
var users = map[string]*User{
	"user1": {
		Password: "password1",
		Role:     "student",
	},
	"user2": {
		Password: "password2",
		Role:     "admin",
	},
}

type User struct {
	Password string
	Role     string
}

// Struct untuk membaca request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Struct claim digunakan sebagai object yang akan di encode oleh jwt
// jwt.StandardClaims ditambahkan sebagai embedded type untuk provide standart claim yang biasanya ada pada JWT
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
