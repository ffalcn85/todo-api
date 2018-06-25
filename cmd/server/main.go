package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ffalcn85/todo-api/internal/memorylists"
)

func main() {
	// Create the todo list service. This can easily be replaced with another service
	// that implements the list.Lists interface if a different implementation is desired.
	transport := httpTransport{memorylists.Create()}

	// Gather the necessary data from commandline flags for the TLS listener and server.
	cert := flag.String("cert", "cert.pem", "The path to the TLS certificate file.")
	key := flag.String("key", "key.pem", "The path to the TLS key file.")
	url := flag.String("url", ":8080", "The address and port for the server to listen on.")
	flag.Parse()

	router := mux.NewRouter()

	router.HandleFunc("/lists", transport.GetListsHandler).Methods(http.MethodGet)
	router.HandleFunc("/lists", transport.AddListHandler).Methods(http.MethodPost)
	router.HandleFunc("/list/{id}", transport.GetListHandler).Methods(http.MethodGet)
	router.HandleFunc("/list/{id}/tasks", transport.AddTaskHandler).Methods(http.MethodPost)
	router.HandleFunc("/list/{id}/task/{taskId}/complete", transport.MarkTaskCompleteHandler).Methods(http.MethodPost)
	err := http.ListenAndServeTLS(*url, *cert, *key, router)
	if err != nil {
		log.Fatal(err)
	}
}
