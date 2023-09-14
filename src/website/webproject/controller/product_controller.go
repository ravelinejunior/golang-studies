package controller

import (
	"html/template"
	"log"
	"module/src/website/webproject/model"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := model.SearchAllProducts()
	temp.ExecuteTemplate(w, "index", products)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "product", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price")
		}

		convertedAmount, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error converting amount")
		}

		model.CreateNewProduct(name, description, convertedPrice, convertedAmount)
	}
	defer Index(w, r) // Redirect to index page after inserting a new product into the database!
	//http.Redirect(w, r, "/", 301) // Redirect to index page after inserting a new product into the database!
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := model.EditProduct(productId)
	temp.ExecuteTemplate(w, "edit", product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		idConvertedToInt, err := strconv.Atoi(id)
		if err != nil {
			panic(err.Error())
		}

		priceConvertedToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price")
		}

		amountConvertedToInt, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error converting amount")
		}

		model.UpdateProduct(idConvertedToInt, amountConvertedToInt, name, description, priceConvertedToFloat)
	}

	defer Index(w, r)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	model.DeleteProduct(productId)

	defer Index(w, r)
	//http.Redirect(w, r, "/", 301)
}
