package main 

import "fmt"


func reformat(message string, formatter func(string) string) string {
// we need to concatinate TEXTIO in the beginnign of the string. 
// add . everytime to the end of the string... 3 times

// Haven't solved full question here, this is the main idea.
	fomatted:= formatter(message + ".")
	fomatted := formatter(message + ".")
	fomatted:= formatter(message + ".")
	return formatter("TEXTIO " + fomatted)

	// Anonymous functions
	func double(a int) int {
    return a + a
}

func main() {
    // using a named function
	newX, newY, newZ := conversions(double, 1, 2, 3)
	// newX is 2, newY is 4, newZ is 6

    // using an anonymous function
	newX, newY, newZ = conversions(func(a int) int {
	    return a * a
	}, 1, 2, 3)
	// newX is 1, newY is 4, newZ is 9
}
}
