package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgryski/dgoogauth"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

var jwtSecret string

type JwtToken struct {
	Token string `json:"token"`
}

type OtpToken struct {
	OTPToken string `json:"otp_token"`
}

func main() {
	router := mux.NewRouter()
	fmt.Println("Starting server at port 8080")
	jwtSecret = "JC7qMMZh4G"
	router.HandleFunc("/signin", SignIn).Methods("POST")
	router.HandleFunc("/verify-otp", VerifyOTP).Methods("POST")
	router.HandleFunc("/protected", ValidateMiddleware(ProtectedEndpoint)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetBearerToken(header string) (string, error) {
	if header == "" {
		return "", fmt.Errorf("an authorization header is required")
	}
	token := strings.Split(header, " ")
	if len(token) != 2 {
		return "", fmt.Errorf("malformed bearer token")
	}
	return token[1], nil
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		bearerToken, err := GetBearerToken(req.Header.Get("authorization"))
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		decodedToken, err := VerifyJwt(bearerToken, jwtSecret)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		if decodedToken["authorized"] == true {
			context.Set(req, "decoded", decodedToken)
			next(w, req)
		} else {
			json.NewEncoder(w).Encode("2FA is required")
		}
	})
}

// mock the user signin
func SignIn(w http.ResponseWriter, req *http.Request) {
	userMock := make(map[string]interface{})
	userMock["username"] = "user1"
	userMock["password"] = "password1"
	userMock["authorized"] = false
	tokenString, err := SignJwt(userMock, jwtSecret)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

func ProtectedEndpoint(w http.ResponseWriter, req *http.Request) {
	decoded := context.Get(req, "decoded")
	json.NewEncoder(w).Encode(decoded)
}

func VerifyOTP(w http.ResponseWriter, req *http.Request) {
	// secret key untuk OTP harus sama antara di google authenticator dan di golang ini
	secret := "2MXGP5X3FVUEK6W4UB2PPODSP2GKYWUT"

	// ambil JWT bearer token dari header
	token, err := GetBearerToken(req.Header.Get("authorization"))
	if err != nil {
		json.NewEncoder(w).Encode("token error")
		return
	}
	decodedToken, err := VerifyJwt(token, jwtSecret)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	//inialisasi config OTP
	otpc := &dgoogauth.OTPConfig{
		Secret:      secret,
		WindowSize:  3,
		HotpCounter: 0,
	}
	var otpToken OtpToken
	_ = json.NewDecoder(req.Body).Decode(&otpToken)

	// validasi OTP menggunakan librarti dgoogauth
	decodedToken["authorized"], err = otpc.Authenticate(otpToken.OTPToken)

	// return invalid juka otp salah
	if decodedToken["authorized"] == false {
		json.NewEncoder(w).Encode("Invalid one-time password")
		return
	}

	// return authorized true pada claim ketika OTP sudah benar
	jwToken, _ := SignJwt(decodedToken, jwtSecret)
	json.NewEncoder(w).Encode(jwToken)
}
