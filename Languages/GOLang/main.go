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

	// Varialbes in Go are passed as values not as references.
	// If we pass a variable to a function, it's value is copied to the function.
	// If we want to pass the reference, we need to pass the address of the variable.
	// We can pass the address of the variable by using the & operator.
	// We can get the value of the address by using the * operator.

	var1 := 10
	increment(var1)
	fmt.Println(var1) // This will print 10, because the value of var1 is not changed. 

	fmt.Println(Getnames())
	
	// ignoring the return value. If function return 2 things but we only want 1 thing. 
	// x, _ := getPoints() // the compiler completely removes from our code. 

	// We can also change the names of the return values. 

	

}

func Getcords(x, y int){
	// This is called naked return statements, just ignore using it. Use it in very simple functions.
	// x and y are already initialized with the default values (0)
	return // this will automatically return x and y (0, 0)
}

// example :- 
// func getPoints(x int, y int) int {
// 	return 3, 4
// }

// functions 
func sum(a, b int) int{ // we can write (x,y int) only when we now that both x and y are integers. 
	return a + b;
}
func sum2(a1 int, b1 int) int { // this is the defalut way. 
	// There is something known as function signatures. If you want to return something from teh function, we need to specify the type of thing which is being returned, in our case it's int. 
	return a1 + b1;
}

func increment(var1 int) int {
	inc := var1 + 1;
	return inc; 
}

// return values concept. If we have more than one return value, we wrap the type of the varialbles like this:- 
func Getnames()(string, string){
	return "Hello", "World";
}
