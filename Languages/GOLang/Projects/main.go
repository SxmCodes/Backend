package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi" 
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("Hello world")
	godotenv.Load(".env") // this will load the .env file in the root directory.	

	portString := os.Getenv("PORT")

	if portString == ""{
		log.Fatal("PORT must be set")
	}
	// with this we will get an error that we need to assign a port everytime we run the program.
	// so we will fetch a package from github := github.com/joho/godotenv
	// this package will help us to read the .env file and assign the values to the variables.
	// we will create a .env file in the root directory and assign the values to the variables.
	fmt.Println("PORT is set to", portString)

	// We will need a router for this, we will use chi chi router.
	// go get -u github.com/go-chi/chi

	router := chi.NewRouter()

	// This is to let users request from their browser. 
	router.Use(cors.Handler(cors.Options{
		// This is basically to allow all the origins to access the server.
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))
	// This project is a json rest api project.

	srv := &http.Server{
		Handler: router,
		Addr: "localhost:" + portString,
	}
	// handling these requests.
	err :=	srv.ListenAndServe()
	if err!=nil {
		log.Fatal(err)
	}
}