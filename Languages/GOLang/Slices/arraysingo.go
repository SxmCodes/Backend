package main

import "fmt"

// in a function, if we want to have an array returned.
func forexample(primary, secondary, tertiary string) ([3]string, [3] string){
	// this means that we will have 2 arrays returned of size 3. 
}

// Varidiac 
func example(nums ...int) int {
	// here nums is just a slice. 
	for i:=0;i<len(nums);i++{
		num := nums[i]
		return num
	}
}

// Spread operator
func example2(strings ...string) {
	
	for i:=0;i<len(strings);i++{
		fmt.Println(strings[i])
	}
}

func main(){
	// Varidiac function calling. - The biggest impact is how we call this function.
	// using ...int is same as using nameSlice [int] but we pass the arguments differently.
	total := example(1,2,3)
	// here I can pass 1,2,3 to the slice (nums). Then the slice can have diff sizes. 
	// fmt.Println is also an example of varidiac function.

	// Spread Operator 
	names := []string {"John", "Doe", "Jane"}
	example2(names...) // here we are spreading the names array. Using spread operator.

	// Thsi is how we declare array in Go 
	var myInts [10]int 
	primes := [6]int{2,3,5,7,11,13}
}
