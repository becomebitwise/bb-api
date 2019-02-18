package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	api "github.com/becomebitwise/bb-api"
	"github.com/go-chi/chi"
)

var salt string

// Serve fires up an HTTP server on the given host/port configuration.
func Serve(host string, port uint, api api.API) error {
	salt = os.Getenv("BB_SALT")
	router := chi.NewRouter()

	router.Get("/", handler)
	router.Post("/login", handleLogin(api.Authenticator()))

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		return err
	}

	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func handleLogin(auth api.Authenticator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds api.Creds
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// TODO (erik): Do some error handling
			log.Println(err)
		}

		if err := json.Unmarshal(body, &creds); err != nil {
			// TODO (erik): Do some error handling
			log.Println(err)
		}

		id, err := auth.Authenticate(r.Context(), creds)
		if err != nil {
			// TODO (erik): Do some error handling
			log.Println(err)
		}

		log.Printf("Authenticated user with ID: %s", id)
		// TODO (erik): Generate token of some kind and give back to client.
		w.Write([]byte("Login successful!"))
	}
}
