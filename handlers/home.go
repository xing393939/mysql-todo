package handlers

import (
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"mysql-todo/libhttp"
	"mysql-todo/models"
	"net/http"
	"strings"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/footer.html", "templates/header.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	brands := make([]models.Brand, 0)
	files, _ := ioutil.ReadDir("./static/brand")
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		brands = append(brands, models.Brand{
			Name: f.Name(),
		})
	}

	data := struct {
		Brands []models.Brand
	}{brands}
	_ = tmpl.Execute(w, data)
}

func GetBrand(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/brand.html", "templates/footer.html", "templates/header.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	brand := mux.Vars(r)["id"]
	products := make([]models.Product, 0)
	files, _ := ioutil.ReadDir("./static/brand/" + brand)
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		products = append(products, models.Product{
			Name: f.Name(),
		})
	}

	data := struct {
		Brand    string
		Products []models.Product
	}{brand, products}
	_ = tmpl.Execute(w, data)
}

func GetBrandProduct(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/brand_product.html", "templates/footer.html", "templates/header.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	brand := mux.Vars(r)["b"]
	productArr := strings.Split(mux.Vars(r)["p"], "-")
	productPre := "/brand/" + brand + "/" + productArr[0] + "/"
	content, _ := ioutil.ReadFile("./static" + productPre + productArr[1] + ".html")
	css1 := productPre + "base.min.css"
	css2 := productPre + productArr[0] + ".css"
	// 替换路径
	contentStr := string(content)
	contentStr = strings.Replace(contentStr, `src="bg`, `src="`+productPre+"bg", 1)

	data := struct {
		Brand   string
		Product string
		Css1    string
		Css2    string
		Content template.HTML
	}{brand, productArr[0], css1, css2, template.HTML(contentStr)}
	_ = tmpl.Execute(w, data)
}

func GetAccountHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tmpl, err := template.ParseFiles("templates/users/dashboard.html.tmpl", "templates/users/home.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	_ = tmpl.Execute(w, nil)
}
