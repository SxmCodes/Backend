package 
import "fmt"

// Creating a matrix. 
func createMatrix (rows, cols int ) [] [] int{
	// length of the slice is 0.
	matrix := make([][]int, 0)
	for i:=0;i<rows;i++{
		// make a new slice to represent a new row.
		row := make([] int, 0) 
		// This inner loop is adding values to each individual row. 
		for j:=0;j<cols;j++{
			row := append(row, i*j)
		}
		// appending all the rows in the matrix. 
		matrix := append(matrix, row) // matrix ke aandar row append karna (in the end).	
	}
	return matrix
}
// Tip:- 
// You must never do this thing:- 
// someSlice := append(otherslice, element) // this is wrong.
// otherslice := append(otherslice, element) // correct. ``


func main(){}