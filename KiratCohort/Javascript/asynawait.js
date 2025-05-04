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