# Packages

So the main package is a stand alone package, it's a special type of package. You will be writing code within this package.

### Library package.

Any other package with a different name other than main. That package is called library package.
It's imported by a main package so that we can use it.

You might know about packages like npm from javscript or pip from python, they are also kind of library package. In which you can code your thing and publish it and other devs can use that thing in their code.

## Naming convention in Packges.

By convention, the package name is same as the last element of it's import path. For example the package math/rand will have files which will begin with

```
package rand
```

## Note

In go, packages lives in directory level not on a file level.

If the code lives in the same directory, it's within the same package. You don't need to import code

Package and directory are the same, you can't have multiple package within same directory.

# Modules

Go programs are organized into packages, A package is a directory of Go code that's all compiled togehter.

A repo contains one or more modules. **A module is a collection of Go packages that are released together**

### Import Paths

If you are familiar with npm or python ecosystem (pip). In Go, there isn't a central location for 3rd party packages, so they all are there in git. **They uses import path as the remote url and download code**

An important point is the standard libarary, you don't need to download standard lib in order to use the code. It comes packaged with go toolchain.

Go use remote URL in go.mod file when we init in that directory. 
- This is basically to download packages. 

Go install let's you to run your code anywhere you want in your whole file system. 

## Making own library package.
You can have your own package name and own functions in it. 
For example :-
```
package onepiece

func Zoro(s string) int {
    best = 1
    return best
}
```
In this example, the function name must be capital that's the main reason that we can export it, if the function first letter of the name is small, then it will be considered as a private function and it won't get exported and we can't use it in our main package. 

It is still useful because Go still checks about compile time errors. 
We can't get .exe files because it's not for production. 

If we want to import that thing in our main program. 
We will use the import path that mtches the module name. example:- github.com/username/mystrings

## Now comes the main part!
In the go.mod file 
we need to replace the github.com url to the package name 

What replace thing is doing is basically it's telling Go that, no need to check the desired package on git or that remote location. To work locally and instead of going and lookin' for it on git just use ..packagename instead.  // local file system. 
```
replace github.com/usrname/package => ../packagename

require(
    that package.
)
```
## NOTE:- 
**This replace thing is good on a local machine or when you are testing something, don't use it on a production env**

**Typically we will push that repo to git and then import it from there.**

**Export as few things as possible when building a library package. Anytime you export any function, you need to suppor that function as well.**

**You need not to change a package exported API, rather you need to focus on to keep changes to internal functionality.**

# Channels 
Go curcurrency ke liye goroutines and channels ka support leta hai. Both of them work togehter and make parallel programmig and sync programming happening all togehter. 

Ab go goroutines aur channels ka funda ye hai ki go routeins uses **channels** for communications between them. 
Channel ka syntax 
```
ch := make(chan int)  // Integer type ka channel
```

Basic channel implementation
```
package main

import "fmt"

func sendData(ch chan int) {
    ch <- 10 // Data send karna
}

func main() {
    ch := make(chan int) // Channel bana rahe hain

    go sendData(ch)  // Goroutine me function call

    data := <-ch  // Channel se data receive karna
    fmt.Println("Received:", data)
}
```
Ab yaha pe hamne channel bna diya of name ch. Ye basically iik pipeline jaise kaam karega, which basically means that ye main function aur go routein mai data bhejne ka kaam karega.

When (ch) functions gets called :-  
Yeh yahan ruk jayega (block ho jayega) jab tak koi receiver nahi hota.
Yani jab tak main() function <-ch use karke value receive nahi karega, sendData() proceed nahi karega.

Same goes while recieving the call, if nothing is being sent in the channel. IT WILL WAIT

This is the flow of Go program
```
Main Goroutine:
   |
   |-----> [Create Channel]
   |
   |-----> [Start Goroutine -> sendData()]
   |                        |
   |                        |-----> [Send 10 to Channel] (BLOCKED)
   |
   |-----> [Receive Data from Channel] (UNBLOCKS sendData)
   |
   |-----> [Print "Received: 10"]
   |
   |-----> [Exit]

```

## Deadlock 
A deadlock is a group of goroutines which are all blocking, so none of them can continue. 

## Closed Channel 
close(ch)

v, ok  := <-ch

You never want to send a value to a closed channel, it will create panic. There is nothing wrong in letting the channel open as it will be garbage collected if never used. 
