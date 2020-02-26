package facade

import (
	"fmt"
	"net/http"
	"os"
)


func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo);
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
	_, _ = fmt.Fprint(w, "Hello, I'm blue")
}

