package main 

import "fmt"

// Maps in Go 
ages := make(map[string] int)
ages["John"] = 37
ages["Mary"] = 24
ages["Mary"] = 21 

/*
	We can perform some functions with maps and keys. 
	1. To delete an element := delete(m, key)
	2. To check if a key is present := value, ok := m[key]
	3. To iterate over a map := for key, value := range m {}
	4. To get the number of elements := len(m)
	5. To create a map with initial values := map[string] int {"John": 37, "Mary": 24}
	6. To create an empty map := map[string] int {}
	7. To create a map with a specific capacity := make(map[string] int, 100)
*/

ages := map[string] int {
	"John": 37,
	"Mary": 24,
	"Noone":39,
}

func main(){
	fmt.Print("Hello World")
}