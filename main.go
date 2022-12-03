package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aanand01762/file-store/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {

	//Create a router and start the server at port 8080
	r := mux.NewRouter()
	routes.RegisterFileStoreRoutes(r)
	fileserver := http.FileServer(http.Dir("./store-files"))
	r.PathPrefix("/").Handler(http.StripPrefix("/", fileserver))
	fmt.Print("Starting web server at 8080\n")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))

}
