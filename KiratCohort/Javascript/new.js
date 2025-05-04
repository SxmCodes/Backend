// Object keys...

function objectMethods(obj) {
  console.log("Original object ", obj);

  // You can think of it like a class and key is the static object of the class Object, input can be any object and it will give you an array which will have keys of that object.
  let keys = Object.keys(obj);
  console.log("After object keys", keys);

  let values = Object.values(obj);
  console.log("After object values", values);

  let entries = Object.entries(obj);
  console.log("After object entries", entries);

  let property = Object.hasOwnProperty("property");
  console.log("After object hasOwnProperty", property);

  // using assign, it will create a new object and copy the properties of the first object to the new object.
  let assign = Object.assign({}, obj, { newproperty: "new value" });
  // Ignore the {}, it basically means that add the newproperty (new value) to the new object and not the original object.
  console.log("After object assign", assign);
}

let obj = {
  name: "John",
  age: 30,
  city: "New York",
};
