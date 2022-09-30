package handlers

import (
	"mysql-todo/models"
	"mysql-todo/libhttp"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/footer.html", "templates/header.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	_ = tmpl.Execute(w, nil)
}

func GetBrand(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/brand.html", "templates/footer.html", "templates/header.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	_ = tmpl.Execute(w, nil)
}

func GetBrandProduct(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/brand_product.html", "templates/footer.html", "templates/header.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	_ = tmpl.Execute(w, nil)
}

func GetAccountHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	sessionStore := r.Context().Value("sessionStore").(sessions.Store)

	session, _ := sessionStore.Get(r, "$GO_BOOTSTRAP_PROJECT_NAME-session")
	currentUser, ok := session.Values["user"].(*models.UserRow)
	if !ok {
		http.Redirect(w, r, "/logout", 302)
		return
	}

	data := struct {
		CurrentUser *models.UserRow
	}{
		currentUser,
	}

	tmpl, err := template.ParseFiles("templates/users/dashboard.html.tmpl", "templates/users/home.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	_ = tmpl.Execute(w, data)
}
