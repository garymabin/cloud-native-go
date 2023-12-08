package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type ProductResponse struct {
	Inventory map[string]int
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/api/echo", echo)
	r.HandleFunc("/api/productAvailability/{productId:[a-zA-Z]+}", checkProduct).Methods("GET")
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
	_, _ = fmt.Fprint(w, "Hello, product service")
}

func checkProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	resp, _ := json.Marshal(&ProductResponse{
		map[string]int{
			params["productId"]: 10,
		},
	})
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	_, _ = fmt.Fprintf(w, "%s", resp)
}
