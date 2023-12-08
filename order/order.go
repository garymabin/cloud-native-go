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
	"strconv"
)

type OrderRequest struct {
	CustomerId string
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/api/echo", echo);
	r.HandleFunc("/api/version", version);
	r.HandleFunc("/api/order/create", createOrder).Methods("POST");
	http.Handle("/", r)
	_ = http.ListenAndServe(port(), nil)
}


func version(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Hello, this is v2 order service")
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
	_, _ = fmt.Fprint(w, "Hello, order service")
}

func paymentService() string {
	paymentService:= os.Getenv("PAYMENT_SERVICE")
	log.Print("paymentService: " +paymentService)
	if len(paymentService) == 0 {
		paymentService = "http://localhost:8080"
	}
	return paymentService + "/api/payment"
}

func makePayment(customerId string, amount int) []byte {
	requestBody, _ := json.Marshal(map[string]string{
		"amount": strconv.Itoa(amount),
	})
	resp, err := http.Post(paymentService()+"/"+customerId, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var orderRequest OrderRequest
	err := decoder.Decode(&orderRequest)

	if err != nil {
		log.Fatal(err)
	}
	response := makePayment(orderRequest.CustomerId, 100)
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "%s", response)
}
