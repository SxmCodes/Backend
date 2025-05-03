// http server 
const http = require("http")
const fs = require("fs")
const url = require("url")


const myserver = http.createServer((req, res) => {
    // creating a log file.
    const log = `${Date.now()}: New request has been recieved!\n`
    const myUrl = url.parse(req.url)
    console.log(myUrl)
    fs.appendFile("log.txt", log, (err, data) => {
        switch (myUrl.pathname) {
            case "/":
                res.end("Homepage")
                break;
            case "/about":
                const username = myUrl.query.myname;
                res.end("About page")
                break;
            default:
                res.end("404 not found...")
        }
    })
});

myserver.listen(8000, () => {
    console.log("Server started!!")
})


