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
		ID:       1,
		Password: "password1",
		Role:     "student",
		City:     "Jakarta",
	},
	"user2": {
		ID:       2,
		Password: "password2",
		Role:     "admin",
		City:     "Bandung",
	},
}

type User struct {
	ID       int
	Password string `json:"-"`
	Role     string
	FullName string
	City     string
}

// Struct untuk membaca request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Struct claim digunakan sebagai object yang akan di encode oleh jwt
// jwt.StandardClaims ditambahkan sebagai embedded type untuk provide standart claim yang biasanya ada pada JWT
type Claims struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
