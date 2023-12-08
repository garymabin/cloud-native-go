package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ProductResponse struct {
	Inventory map[string]int
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/api/echo", echo);
	r.HandleFunc("/api/shoppingCart/purchase", purchaseCart).Methods("POST");
	http.Handle("/", r)
	_ = http.ListenAndServe(port(), nil)
}

func productService() string {
	productService:= os.Getenv("PRODUCT_SERVICE")
	log.Print("productService: " + productService)
	if len(productService) == 0 {
		productService = "http://localhost:8080"
	}
	return productService + "/api"
}

func orderService() string {
	orderService:= os.Getenv("ORDER_SERVICE")
	log.Print("orderService: " + orderService)
	if len(orderService) == 0 {
		orderService = "http://localhost:8080"
	}
	return orderService + "/api/order"
}

func purchaseCart(w http.ResponseWriter, r *http.Request) {
	shoppingcart := [2]string{"productA", "productB"}
	productAvailabilityServiceUrl := productService() + "/productAvailability/"
	for i := range shoppingcart {
		resp, err := http.Get(productAvailabilityServiceUrl + shoppingcart[i])
		if err != nil {
			log.Fatal(err)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		var productResponse ProductResponse
		_ = json.Unmarshal(body, &productResponse)
		log.Printf("productResponse: %s", body)
	}
	orderRequestBody, _ := json.Marshal(map[string]string{
		"customerId": "abcdef",
	})
	resp, err := http.Post(orderService()+"/create", "application/json", bytes.NewBuffer(orderRequestBody))
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "%s", body)
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]
	w.Header().Add("Content-Type", "text/plain")
	_, _ = fmt.Fprint(w, message)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8081"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Hello, I'm blue")
}

