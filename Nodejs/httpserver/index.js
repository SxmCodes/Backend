// http server 
const http = require("http")
const fs = require("fs")

const myserver = http.createServer((req, res) => {
    // creating a log file.
    const log = `${Date.now()}: New request has been recieved!\n`
    fs.appendFile("log.txt", log, (err, data) => {
        console.log(req.headers)
        res.end("Hello from my new server!!")
    })
});

myserver.listen(8000, () => {
    console.log("Server started!!")
})


