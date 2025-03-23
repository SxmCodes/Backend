package main 

import (
	"fmt"
	"os"
	"log"
	
	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("Hello world")

	portString := os.Getenv("PORT")

	godotenv.Load(".env") // this will load the .env file in the root directory.	

	if portString == ""{
		log.Fatal("PORT must be set")
	}
	// with this we will get an error that we need to assign a port everytime we run the program.
	// so we will fetch a package from github := github.com/joho/godotenv
	// this package will help us to read the .env file and assign the values to the variables.
	// we will create a .env file in the root directory and assign the values to the variables.
	fmt.Println("PORT is set to", portString)
}