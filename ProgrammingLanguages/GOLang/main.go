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

	// Declaring multiple variables at once.
	type1, type2 := "Hello world 2", "Hello world 3";
	fmt.Println(type1)
	fmt.Println(type2)

	// Function call.
	fmt.Println(sum(1,2));
	fmt.Println(sum2(1,20));

	
}

// functions 
func sum(a, b int) int{ // we can write (x,y int) only when we now that both x and y are integers. 
	return a + b;
}
func sum2(a1 int, b1 int) int { // this is the defalut way. 
	// There is something known as function signatures. If you want to return something from teh function, we need to specify the type of thing which is being returned, in our case it's int. 
	return a1 + b1;
}