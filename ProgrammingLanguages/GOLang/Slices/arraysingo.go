package main

// in a function, if we want to have an array returned. 
func forexample(primary, secondary, tertiary string) ([3]string, [3] string){
	// this means that we will have 2 arrays returned of size 3. 
}

/*          ARRAYS AND SLICES. 
*      Arrays are what we know of
*      Slices are build on top of array, somewhat a mini array you can say. 
*      For example :-
*      	If I have array a:= [1,2,3,4,5]
*      	I can have a slice of this array like b := a[1:4] - 4th index not included.
*     NOTE:- 
*    	 If you are using slices inside a function. You have to use sampleslice[:] - This means that you are returning every index of that array.
* */

func main(){
	// Thsi is how we declare array in Go 
	var myInts [10]int 
	primes := [6]int{2,3,5,7,11,13}
}
