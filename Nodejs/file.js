// In order to interact with files. We need to import fs module.
const fs = require("fs");
    
// fs.readFileSync() - Read a file synchronously
fs.writeFileSync("./hello.txt", "Hello World!"); // Create a file and write to it 
const data = fs.readFileSync("./hello.txt", "utf-8"); // Read the file synchronously

// Async will not return you anything. If you put this in a variable, it will be undefined.
fs.readFile("./hello.txt", "utf-8", (err, result) => { 
    if (err) {
        console.log(err);
        return;
    }
    console.log(result);
}); 