function nothing() {
    let promise = new promise(function (resolve) {
        setTimeout(() => {
            console.log('Hi there');
        }, 2000);
    });
    return promise;
}
  
// instead of using .then() and .catch() we can use async await
async function main() {
    // when you do this, it's the same as saying...
    let value = await nothing()

    // This is also same...
    nothing().then(function (value) {
        console.log(value);
    })

    console.log(value);
}
// usually you don't write your own async function, you just call it, and wrap it around another async function. 

// Today's assignment is to write own async function which does a callback to the thread of cpu. How to get into main thread and ask js to make raw async function. 
// Next assignment is to write your own promise class, which will have resolve and duration. 