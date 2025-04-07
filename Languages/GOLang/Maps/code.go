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
	2. To check if a key is present := value, ok := m[key] - here ok is boolean, if boolean is true then key is present.
	3. To iterate over a map := for key, value := range m {}
	4. To get the number of elements := len(m)
	5. To create a map with initial values := map[string] int {"John": 37, "Mary": 24}
	6. To create an empty map := map[string] int {}
	7. To create a map with a specific capacity := make(map[string] int, 100)
*/

// name will take the username which we need to detete. Users is a map of string (usernames) and user is a structure. 
func deleteIfNecessary(users map[string] user, name string ) (deleted bool, err error ){
	// deleted variable will tell weather given user is deleted or not. Error aaigi toh return karega (err variable)
	// check weather the user is present in the map or not. If not present then return false and error.
	if _, ok := users[name]; !ok {
		// errorf basically iik error object banata hai formatted string ke sath. Ye errors handling mai use hota hai. Hm iske sath customized erros bna sakte hai and values bhi format kar sakte hai. Also error wrapping mai bhi kaam aata hai.

		return false, fmt.Errorf("User %s not found", name)
	}
	delete(users, name)
	return true, nil
}

// This is the sample function which tells us how to use error handling using fmt.Errorf
func divide(a, b int) (int, error){
	if b == 0{
		fmt.Errorf("Cannot divide by zero")
		return 0, err
	}
	return a/b, nil
}
// also some sample functions to know about advance functions. // The first class functions. 
func add (a, b int) int {
	return a+b
}
func mult (a, b int) int {
	return a*b
}

func aggregate(a, b, c int, arth func(int, int) int) int {
	return arth(arth(a,b), c)
}

ages := map[string] int {
	"John": 37,
	"Mary": 24,
	"Noone":39,
}

func main(){
	fmt.Print("Hello World")
}