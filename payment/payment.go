package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

type PaymentRequest struct {
	Amount string
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/api/echo", echo)
	r.HandleFunc("/api/payment/{customerId:[a-z]+}", createPayment).Methods("POST")
	http.Handle("/", r)
	_ = http.ListenAndServe(port(), nil)
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]
	w.Header().Add("Content-Type", "text/plain")
	_, _ = fmt.Fprint(w, message)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Hello, payment service")
}

func createPayment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	body, _ := ioutil.ReadAll(r.Body)

	var paymentRequest PaymentRequest
	err := json.Unmarshal(body, &paymentRequest)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "%s$ payment for %s is now processing", paymentRequest.Amount, params["customerId"])
}
