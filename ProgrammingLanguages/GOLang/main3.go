package main

import (
	"fmt"
)

func main() {
  var a, b = 6, "Hello"
  c, d := 7, "World!"

  // we can assign multiple varialbes like this:- 
  var (
      yoo int
      yoo1 int = 1
      yoo2 string = "hello"
   )

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)

  fmt.Println(yoo)
  fmt.Println(yoo1)
  fmt.Println(yoo2)
}