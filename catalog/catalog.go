package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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
	productsUrl = os.Getenv("PRODUCT_URL")
}

func loadProducts() []Products {
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
	loadProducts()
}
