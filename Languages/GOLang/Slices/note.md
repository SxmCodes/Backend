## ARRAYS AND SLICES. 
Arrays are what we know of
Slices are build on top of arraysomewhat a mini array you can say. 

For example :-
If I have array a:= [1,2,3,4,5]
I can have a slice of this arraylike b:= a[1:4] - 4th index notincluded.
### Note:- 
   	 If you are using slices inside a function. You have to use sampleslice*[:] - This means that you are returning every index of that array.

#### Make 
This is used when we neeed to make a slice but we don't need to specify an array for that. 

Syntax is:- 
```
mySlice := make([]int, 5, 10)
```
This basically means that It will create an array under the hood, and we will have the slice of 5 indexes and 10 will be the size of the array. Now you basically don't even need to specify the length of the array. 

You can have 
```
mySlice := make([]int, 5)
```

We also have build in functions, len() and cap() functions. 

## Variadic 
This is the way to initilize a slice in function. 
Example with code is mentioned in arraysingo.go file. 

## Spread
Spread operator is inverse of varidiac function. 
We can take a piece of slice and pass it into the variadic function. 

## Append 
If the underlying array is not large enough, append() will create a new underlying array and point the returned slice to it. Append is also varidiac.

Everything will work here!
```
slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)
```

## Range 
Go me range ek shortcut hai loops likhne ke liye. Ye slice (ya array) ke har element pe iterate karta hai aur uska index aur value ek sath deta hai.

```
for INDEX, ELEMENT := range SLICE {
    // kuch bhi kar sakte ho yaha
}
```
Index:- Will tell the index. 
Element:- Will tell the actual data on that index. 

Sample Go code for example :- 

```
fruits := []string{"apple", "banana", "grape"}

// similar hai python se bss in nahi laga lol. 
// i - storing index, fruit - storing the acutal name of the fruit.
for i, fruit := range fruits {
    fmt.Println(i, fruit)
}

// ignoring the index :- 
for _, fruit := range fruits {
    fmt.Println(i, fruit)
}
```

## Slices of Slices. 
This is basically a matrix. 2D matrix. 

```
rows := [][] int{}
```

