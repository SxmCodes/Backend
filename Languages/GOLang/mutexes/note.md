# Mutexes in Go

This will allow us to lock access to data. This ensures that we can control which goroutines can access certain data at which time.
This allows us to share data between multiple goroutines.

Go standard lib provides implementatoin with mutexes with sync.Mutex 1. .lock 2. .Unlock

🛠 What is a Mutex?
A mutex (mutual exclusion) is a locking mechanism used to ensure that only one goroutine can access a critical section of code at a time. This prevents race conditions in concurrent programs.

Go provides a built-in sync.Mutex type with two primary methods:

Lock() - Locks the mutex, blocking other goroutines from accessing the same resource.
Unlock() - Unlocks the mutex, allowing other goroutines to proceed

```
func protected(){
    mux.Lock()
    defer mux.Unlock()
}
```

### NOTE:-

    We use defer in front of unblock to ensure that there would be no errors and that would be unlocked after whole function has been executed. It basically ensures that there are no deadlocks!

# RWMutex

🛠 What is sync.RWMutex?
sync.RWMutex (Read-Write Mutex) is an optimized version of sync.Mutex that allows:

Multiple readers (RLock) at the same time ✅
Only one writer (Lock) at a time 🚨
Matlab:

Agar sirf read karna hai → Multiple goroutines ek saath read kar sakti hain (RLock()).
Agar write karna hai → Ek time pe sirf ek goroutine access karegi (Lock()), aur tab koi bhi read nahi kar sakta.

## 📌 Methods of `sync.RWMutex`

| Method      | Description                                               |
| ----------- | --------------------------------------------------------- |
| `Lock()`    | Acquires exclusive lock (blocks all readers and writers). |
| `Unlock()`  | Releases exclusive lock.                                  |
| `RLock()`   | Acquires shared lock (multiple readers allowed).          |
| `RUnlock()` | Releases shared lock.                                     |

📌 Why Use RWMutex?
Faster than Mutex for read-heavy operations.
Readers don’t block each other, but they do block writers

📌 Key Takeaways
✅ Use RWMutex when reads > writes (read-heavy workloads).
✅ Multiple RLock() calls don’t block each other.
✅ Lock() blocks everyone (readers & writers).
✅ Improves performance compared to sync.Mutex.

# Generics in Go

enerics allow programmers to write behavior where **the type can be specified later because the type isn’t immediately relevant.** This is an amazing feature because it allows us to write abstract functions that drastically reduce code duplication.

For example the following code will split up slices in half no matter the type of slice. This is before generics btw

```
func splitSlice(s [string]) [string] [string]{
    mid := len(s)/2
    return s[:mid], s[mid:]
}
```

### After Generics

Here everything is type safe, that we know what we are giving (type of slice), we will get back same type.
It is to make our code DRY - Don't repeat yourself.

```
func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
```

Here T is the name of the type parameter, which can be anything and we didn't have to specify the type again and again. This helps preventing code duplication.

# Constraints

Constraints are just **interfaces that allow us to write generics that only operate within the constraints** of a given interface type. In the example above, the any constraint is the same as the empty interface because it means the type in question can be anything.

In this code, we are passing stringer to function, type of vals is stringer (which has string attribute on it). We were able to pass it and use it inside the function.

```
type stringer interface {
    String() string
}

func concat[T stringer](vals []T) string {
    result := ""
    for _, val := range vals {
        // this is where the .String() method
        // is used. That's why we need a more specific
        // constraint instead of the any constraint
        result += val.String()
    }
    return result
}
```
