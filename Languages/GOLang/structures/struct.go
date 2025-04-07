import "fmt"
package main 
// Anonymous Structures - don't have a name.Mycar is not the name of the sturct, it's the instance of the struct. 
Mycar := struct{
	Make string 
	Model string
}{
	Make: "BMW",
	Model: "Model A"
}

// You can basically have Anonymous strucutes within other fields
type car struct {
	Make string
	Model string

	// we can make other fields as well. It contains Anonymous Structures. It will always be used one time.  
	Wheel, struct {
	Radius int
	Material int
	}
}

// Embedded Structures. 
type car struct {
	make string
	model string
}

type truck struct {
	car  //  we have embedded struct car in truck. If we need to access some parameter in car, for example we need to have radius of the wheel, We don't need to do truck.car.radius, we can directly do truck.radius.
	// we are inheriting all the parameters from car type. 
	bedsize int 
}

func main(){

}
