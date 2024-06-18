package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	ID    int     `json: "id"`
	Name  string  `json: "name"`
	Price float64 `json: "price"`
}

var products []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//for _, item := range products {
	//	if item.ID, _ := strconv.Atoi(params["id"]); item.ID == item.ID
	//	{json.NewEncoder(w).Encode(item)
	//	return
	//	}
	//	}
	//	json.NewEncoder(w).Encode(&Product{})
	//}

	// Check if "id" parameter exists
	if id, ok := params["id"]; ok { // Added opening curly brace

		// Convert "id" to integer and handle potential errors
		itemID, err := strconv.Atoi(id)
		if err != nil {
			// Handle error: invalid id format, return appropriate error repo
			http.Error(w, "Invalid product ID format", http.StatusBadRequest)
			return
		}

		// Find product by ID
		for _, item := range products {
			if item.ID == itemID {
				json.NewEncoder(w).Encode(item)
				return
			}
		}

		// No product found with the given ID
		json.NewEncoder(w).Encode(&Product{}) // Assuming Product is a struct defining a product
		return
	} else {
		// Handle missing "id" parameter
		http.Error(w, "Missing required parameter 'id'", http.StatusBadRequest)
		return
	}
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}
func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Check if "id" parameter exists
	if id, ok := params["id"]; ok {

		// Convert "id" to integer and handle errors
		_, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid product ID format", http.StatusBadRequest)
			return
		}

		// Decode request body into a Product struct
		var updatedProduct Product
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&updatedProduct)
		if err != nil {
			http.Error(w, "Invalid product data in request body", http.StatusBadRequest)
			return
		}

		// Validate updated product data (optional)
		// You can add checks here to ensure required fields are present, etc.

		// Update product in your data store (replace with your logic)

		// Respond with the updated product or success message

		json.NewEncoder(w).Encode(updatedProduct)

	} else {
		http.Error(w, "Missing required parameter 'id'", http.StatusBadRequest)
		return
	}
}
func main() {
	products = append(products, Product{6587015, "Coke", 12.00})
	products = append(products, Product{6587014, "Water", 7.25})
	r := mux.NewRouter()
	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/api/products", createProduct).Methods("POST")
	r.HandleFunc("/api/products/{id}", updateProduct).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", r))

}
