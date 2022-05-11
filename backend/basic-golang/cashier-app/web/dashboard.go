package web

import "net/http"

func (web *WEB) DashboardPage(w http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{
		"title": "Dashboard",
		"name":  "Dashboard Page",
	}
	tmpl, err := web.BasePath()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "dashboard", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
