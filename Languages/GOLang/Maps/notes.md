# Maps in Go 
Any type can beocome a value in maps but only few types can become keys.  

## Key types 
Any type can be used as the value in a map, but keys are more restrictive. 

What does a type qualify in order to become a map key?
A type must be comparable like strings, bool, numbers can be comparable

 functions can't be comparable, they are just pointers who addresses an addresses in memory. 

Slices can be compared and can be treated as a key. 

If you are comparing one slice to another slice, you are not comparing the underlined values, you are comparing where those 2 slices are stored in your RAM.

# Nested maps
Maps ke aandar maps. 
These are not much used. 

Using a structure inside a map as a key is simpler to do. 

## Runes
Runes are individual charaters, it is an alias of int32 and it's used to represents charcters.
Example :- 
```
package main

import "fmt"

func main() {
    ch := 'A'  // Single quotes me diya to rune ban gaya
    fmt.Println(ch)      // Output: 65 (Unicode value of 'A')
    fmt.Printf("%c\n", ch) // Output: A
}
```
