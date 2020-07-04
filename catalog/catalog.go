package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type Product struct {
	Uuid 		string `json:"uuid"`
	Product string `json:"product"`
	Price 	float64 `json:"price,string"`
}

type Products struct {
	Products []Product
}

var productsUrl string

func init() {
	// Example: export PRODUCT_URL=http://localhost:8081
	productsUrl = os.Getenv("PRODUCT_URL")
}

func loadProducts() []Product {
	response, err := http.Get(productsUrl + "/products")
	if err != nil {
		fmt.Println("Erro de HTTP")
	}
	data, _ := ioutil.ReadAll(response.Body)

	var products Products
	json.Unmarshal(data, &products)

	fmt.Println(string(data))
	return products.Products
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", ListProducts)
	r.HandleFunc("/product/{id}" , showProduct)
	http.ListenAndServe(":8080", r)
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadProducts()
	t := template.Must(template.ParseFiles("templates/catalog.html"))
	t.Execute(w, products)
}

func showProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response, err := http.Get(productsUrl + "/product/" + vars["id"])
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)

	var product Product
	json.Unmarshal(data, &product)

	t := template.Must(template.ParseFiles("templates/view.html"))
	t.Execute(w, product)
}
