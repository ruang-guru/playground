// |=======>>> Sessions <<<=======|
// Cara menyimpan data session cookies menggunakan gorilla/sessions package yang populer di Go.

// Cookie adalah bagian kecil dari data yang disimpan di browser user dan dikirim ke server pada setiap request.
// Di dalamnya, kita dapat menyimpan misalnya, apakah pengguna masuk ke situs web atau tidak dan mencari tahu siapa dia sebenarnya (di sistem).
// Dalam contoh ini kita hanya mengizinkan user yang diautentikasi untuk melihat pesan rahasia di halaman /secret.
// Untuk mengaksesnya, pertama-tama dia harus mengunjungi /login untuk mendapatkan session cookie yang valid, yang membuatnya login.
// Selain itu, dia dapat mengunjungi /logout untuk mencabut aksesnya ke secret message.

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}

// $ curl -s http://localhost:8080/secret
// Forbidden

// $ curl -s -I http://localhost:8080/login
// Set-Cookie: cookie-name=MTQ4NzE5Mz...

// $ curl -s --cookie "cookie-name=MTQ4NzE5Mz..." http://localhost:8080/secret
// The cake is a lie!
