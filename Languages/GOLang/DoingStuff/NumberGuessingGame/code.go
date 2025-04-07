package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	// rand.intn returns a random int n, 0 <= n < 100 and rand.Float64 returns a random float64, 0.0 <= f < 1.0

	// seed the random number generator with the current time
	// concept of seeding and seed:- 
	/*
	Computers don't produce truly random numbers. Instead, they generate sequences of numbers that appear random but are actually determined by an algorithm. These are called pseudo-random numbers.
	SEED := The "seed" is the starting value for this algorithm. Think of it as the initial input that dictates the entire sequence of "random" numbers that will follow.
	If you use the same seed, you'll get the same sequence of numbers every time. If you use a different seed, you'll get a different sequence.
	*/

	// this is deprecated in go 1.20
	// rand.Seed(time.Now().UnixNano()) // This is to speed up the process. Seed is the current time which is converted into int64. 

	// this gives constantly changing number.
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src) 

	// gettin a random number between 0 and 100
	target := rng.Intn(100) +1 // +1 to avoid 0
	var guess, attempts int

	fmt.Println("Guess a number between 1 and 100")

	for{
		// take user input.
		fmt.Print("Make a guess: ")
		 _, err:= fmt.Scan("%d", &guess)
		// input validation
		 if err!=nil {
			fmt.Println("Invalid input.")
			continue
		 }
		attempts++

		if guess < target{
			fmt.Println("Too low")
		}else if guess > target{
			fmt.Println("Too high")
		}else{
			fmt.Println("You got it!")
			break
		}
	}
}