// It will still use callback queue, event looops exitCode.
const cart = ["shoes", "pants", "socks"];

function proceedtopayement(orderid) {
  console.log("Payment is being processed for order id: " + orderid);
}
// Why this is bad is because, here we gave the whole function into createOrder API, we were relying on the fact that createOrder will call the function after it is done with the order creation. Which is not a good practice.
createOrder(cart, function (orderid) {
  proceedtopayement(orderid);
});

//  using promises. This is more readable and maintainable code.
// In here we are attaching a callback to the promise. The promise will call the callback when it is done with the order creation.
const promise = createOrder(cart);
promise.then(function (orderid) {
  return proceedtopayement(orderid);
});

// There is difference between passing a function and attaching a function into promise object.
// So the promise will call the function once the function has the data inside it. It will call it just once.

// This is the api given by browsers to us in order to make external calls.
const GITHUB_API = "https://api.github.com/users/SxmCodes";

// This will return a promise.
const user = fetch(GITHUB_API);
console.log(user);

user.then(function (data) {
  console.log(data);
});

// instead of promise.then we can also write...
createorder(cart).then(function (orderid) {
  console.log("Order id is: " + orderid);
});
// this is equavalent to writing it as like on line 14th. Now we can use as much .then() as we want. That will eventually becomes that we call promise chaining.

// one great tip here is that we always return the promise from the function. So that we can chain it with other promises. This is a good practice. So this is one example to summarise the topic...
createorder(cart)
  .then(function (orderid) {
    console.log("Order id is: " + orderid);
    return proceedtopayement(orderid);
  })
  .then(function (paymentid) {
    return console.log("Payment id is: " + paymentid);
  })
  .then(function (orderid) {
    console.log("Order id is: " + orderid);
    return proceedtopayement(orderid);
  })
  .then(function (paymentid) {
    return console.log("Payment id is: " + paymentid);
  });
