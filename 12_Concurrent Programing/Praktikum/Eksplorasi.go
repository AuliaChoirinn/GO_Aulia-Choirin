package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
}

func fetchProducts(url string, products chan<- []Product) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var productsData []Product
	err = json.Unmarshal(body, &productsData)
	if err != nil {
		fmt.Println("Error decoding response body:", err)
		return
	}

	products <- productsData
}

func main() {
	productsCh := make(chan []Product)
	url := "https://fakestoreapi.com/products"

	go fetchProducts(url, productsCh)

	products := <-productsCh

	for _, p := range products {
		fmt.Printf("%d\t%s\t%.2f\n", p.ID, p.Title, p.Price)
	}
}
