## ARRAYS AND SLICES. 
*      Arrays are what we know of
*      Slices are build on top of array, somewhat a mini array you can say. 
*      For example :-
*      	If I have array a:= [1,2,3,4,5]
*      	I can have a slice of this array like b := a[1:4] - 4th index not included.
### Note:- 
*    	 If you are using slices inside a function. You have to use sampleslice*[:] - This means that you are returning every index of that array.

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
