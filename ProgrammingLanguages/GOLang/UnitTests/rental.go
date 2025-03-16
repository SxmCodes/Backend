package main 

import "fmt"

package main

// Named return parameters are great for documentation, we know what function is returning. Named Return parameters are important for longer functions which have many return values. 
// Named return parameters are also useful for functions that have multiple return values.

// Returning early from a function. 
// Guard Clauses leverage the ability to return early from a function (or continue through a loop) to make nested conditionals one-dimensional. Instead of using if/else chains, we just return early from the function at the end of each conditional block.



func yearsUntilEvents(age int) (
	yearsUntilAdult int, 
	yearsUntilDrinking int, 
	yearsUntilCarRental int)
	 {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return 
}

