function nothing() {
  let promise = new promise(function (resolve) {
    setTimeout(() => {
      console.log("Hi there");
    }, 2000);
  });
  return promise;
}

// instead of using .then() and .catch() we can use async await
async function main() {
  // when you do this, it's the same as saying...
  let value = await nothing();

  // This is also same...
  nothing().then(function (value) {
    console.log(value);
  });

  console.log(value);
}
// usually you don't write your own async function, you just call it, and wrap it around another async function.

// Today's assignment is to write own async function which does a callback to the thread of cpu. How to get into main thread and ask js to make raw async function.
// Next assignment is to write your own promise class, which will have resolve and duration.

// here is one mst example...
function myownfunction() {
  return new Promise(function (onDone, onError) {
    onError("Error");
  });
}

const p = myownfunction();

p.then(function () {
  console.log("Success");
}).catch(function () {
  console.log("Error");
});

// one more example 
function arthematic(a, b, fn) {
  let value1 = a * a; 
  let value2 = b * b;
  fn(value1, value2);
}

arthematic(2, 3, function (value1, value2) {
  console.log(value1 + value2);
}
);

// returning as a promise. 
arthematic(1, 2).then(function (value) {
console.log(value);
})