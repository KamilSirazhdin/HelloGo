package main
// git 
import (
	"fmt"
	"log"
	"net/http"
	"github.com/rs/cors"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"text": "Hello Go!"}`)
}

func main() {
	http.HandleFunc("/api/hello", handler)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	})

	handler := c.Handler(http.DefaultServeMux)
	log.Println("Starting server on :5000")
	log.Fatal(http.ListenAndServe(":5000", handler))
}
