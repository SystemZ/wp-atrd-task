package rest

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func StartWebServer() {
	// define all REST API routes
	r := mux.NewRouter()
	r.Handle("/v1/secret", http.HandlerFunc(NewSecret)).Methods("POST")
	r.Handle("/v1/secret/{hash}", http.HandlerFunc(GetSecret)).Methods("GET")
	r.Handle("/", http.HandlerFunc(GetDocsRedirect)).Methods("GET")
	r.Handle("/v1/docs/", http.HandlerFunc(GetSwaggerUi)).Methods("GET")
	r.Handle("/v1/swagger.yml", http.HandlerFunc(GetSwaggerYml)).Methods("GET")

	// add logging to our router
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	// start HTTP server
	logrus.Fatal(http.ListenAndServe(":3000", loggedRouter))
}
