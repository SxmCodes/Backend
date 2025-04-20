# Error handling in Go 

In Go errors are not handled as they are hendled in JS

Error handled in JS, have try catch situation
```
JS

function main(){
try {
const user = getuser()
const profile  = getuserProfile(user.ID)
} catch (err){
c.log(err)
}
}
```

and in GO 
```
package main

import "fmt"

func main(){
user, err := getUser()
if err!=nil {
fmt.println(err)
return 
}

// for example, we need to fetch the user profile. If there is any error with profile it will print that, if not then we have the profile. 
profile, err := getuserProfile(user.ID)
if err!=nil {
fmt.Println(err)
return 
}
}
```

As I did with the profile in Go, I basically can't do that thing in Javscript. I can't copy paste another try catch thing for user profile. 
**This is because the varialbe user is only avaliable for that first try catch block, because of it's scope.**

To do the error handling of the profile thing, we need to write that thing in that same try block 

```
try {
const user = getUser()
const profile = getuserProfile(user.ID)
} catch (err){
c.log(err)
}
```
With this implementation, all the errors are in one single place. If we want another profile function error, we need to have nested try catch blocks, for individual errors.  

We can' really know weather the function is dangerous (can get error) or it's safe.

We can also explictly mention the error in functions signature, but in JS we need to see functions in order to tell that the specific function is returning an error.


#### Error in Go are just values and specifically they are just interfaces. 

The built in error interface is the interaface with just one method :- the error method
example:- 
```
type error interface {
Error() string
}
```
and it will return a string describing what went wrong. 
