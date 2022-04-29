package web

import "net/http"

func (web *WEB) loginPage(w http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{
		"title": "Sign In",
	}
	tmpl, err := web.BasePath()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "login", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (web *WEB) registerPage(w http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{
		"title": "Sign Up",
	}
	tmpl, err := web.BasePath()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "signup", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
