# Modules bakchodi... 

If you have math.js file (module) in your directory and you need to call that shit in your main.js 

You need to do this thing := 
```
const math = require(./math.js)
```

then you can use anything that is present inside the math.js file (module)

Here there are 2 things to worry about:-

If you type this shit... 
```
const math = require(./math.js)

console.log(math)
```
this file will return {}, an empty object. 

#####  You need to export that function which you have created.

If you export anything from a file, the function becomes that shit.
```
this is from math.js file

function add(a, b){
    return a+b;
}

module.exports("Iron man")

// If we do this then we will get desired result:-
module.exports(add)
```

now the function is exported as iron man, so if we run that shit in main.js so we will get Iron man instead of {} an empty object. 

## Exports... 

You can use this thing in order to export any function. 
```
module.exports= {
add, 
fun2
}
```

This is to export while using export function. add() and subs() are property, but they are anonymous function because they have no name...
```
exports.add() = (a, b) => a+b;

exports.subs() = (a,b) => a-b;
```
You can use module.exports only one time, because it will overwrite the value. 
