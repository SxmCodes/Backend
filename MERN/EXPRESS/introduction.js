// This is the introduction page of the express js 
const express = require('express'); // it imports the libraries of express and stores it in express variable.
const app = express(); // we are creating a new express app, here the varlialbe app is the main component of the application in which we will fire middlewares, routes that will be used to handle the requests and responses.
const port = 3000;

app.set('view engine', 'ejs'); // Setting up EJS as the view engine, EJS means Embedded JavaScript, it is a templating engine that allows us to generate HTML markup with plain JavaScript. Server side javscript ko html mai mix kar sakte hai. 

// starting the server. 
app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});

// defining the route for the home page.
app.get('/', (req, res) => {
  res.render('index', { title: 'Welcome to My Website', message: 'Hello, This is my first express application!' }); // res.render is a method of express which is used to render views, in this case it is rendering the index.ejs file.
  res.send('Hello World!'); // res.send is a method of express which is used to send the response to the client.
});