package main

import (
	"fmt"
	"crypto/sha256"
	
)

// handler function.
// func returnHandler(w http.ResponseWriter, r *http.Request){
	
// }	


func main(){
	
	s := "This is the new age of AI"
	h := sha256.New()
	h.Write([]byte(s)) // will create slice of byte and then writes it into hash instance. 
	nothing := h.Sum(nil) // returns the hash value in byte slice. Nil means that we are not appending anything to result.
	fmt.Printf("The SHA256 hash of '%s' is: %x\n", s, nothing) // %x is used to print the byte slice in hexadecimal format.
}