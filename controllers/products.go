package controllers

import (
	"log"
	"main/models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", models.GetAllProducts())
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConvertidoParaFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantityConvertidaParaInt, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("Erro na conversão do quantity:", err)
		}

		models.CreateNewProduct(name, description, priceConvertidoParaFloat, quantityConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}
