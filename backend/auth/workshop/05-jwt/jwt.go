package main

import "github.com/dgrijalva/jwt-go"

// Jwt key yang akan dipakai untuk membuat signature
var jwtKey = []byte("key")

// Data user - password yang bisa mengakses api
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
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
	jwt.StandardClaims
}
