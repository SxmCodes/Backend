package main

import ("fmt")

func Append(slice, data []byte) []byte {
	l := len(slice)
	if l + len(data) > cap(slice) { // capacity of total slice. 
		// Need to grow the slice
		newSlice := make([]byte, (l+len(data))*2) // allocating double, because we will not reinitialize the slice again and again..
		// The copy function is used to copy the data from one slice to another
		copy(newSlice, slice)
		slice = newSlice // swapping slice with newSlice. 
	}
	// appending the new data. 
	slice = slice[0:l+len(data)]
	copy(slice[l:], data) // jitna purana data tha uske baad se naye data ko chipka do with the help of copy function.
	return slice 
}

func main(){

}