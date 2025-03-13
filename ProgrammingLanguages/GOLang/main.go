// This tells the compiler that this is the main starting point of the program
package main

// Importing the fmt package from the standard library, which has all the necessary functionality.
import (
	"fmt"
);

func main(){
	// Declare and initialize a variable. 
	var yoo1 string = "Hello world 0";
	fmt.Println(yoo1)
	// This is the other way to declare a variable, instead or mentioning about the datatype. 
	yo := "Hello world 1";
	fmt.Println("Hello world again")
	fmt.Println(yo)
}