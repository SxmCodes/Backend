 ## Imp functions.
 handle function will handle anything which comes under /. It writes response back to the client, and *http.request contains the info about the incoming requests. 
	 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	 	fmt.Fprintf(w, "Hello!")
	 })
	
	 redirecting function. 
	 http.HandleFunc("/abc123", func(a http.ResponseWriter, c *http.Request){
		 Redirect replies to the request with a redirect to url, which may be a path relative to the request path.
		 StatusMovedPermanently is the status code for Moved Permanently responses.
	 	http.Redirect(a, c, "/", http.StatusMovedPermanently)
	 })
	 ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
	 ListenAndServe always returns a non-nil error.
	 The handler is typically nil, in which case the DefaultServeMux is used.
	 http.ListenAndServe(":8080", nil)
	
## Imp concepts
### Post endpoint.
In the context of web development and APIs (Application Programming Interfaces), a "POST Endpoint" refers to a specific URL (Uniform Resource Locator) that is designed to receive data from a client (like a web browser or mobile app) and then process that data on the server. Here's a breakdown: It's the "doorway" through which a client can interact with the server's functionality.

**It contains mostly HTTP**

### HTTP Request Validation.
An HTTP request validation pattern is a security and design practice used in web applications to ensure that incoming HTTP requests meet specific criteria before processing them. This pattern helps prevent unauthorized or incorrect access to server-side resources and protects against potential misuse of API endpoints.

### Input sanitization.
Input sanitization is the process of modifying or removing potentially harmful data entered by users to prevent web-based attacks like SQL injection and cross-site scripting (XSS). 

## What code is doing.
r.Method represents the **HTTP method of the incoming request** (GET, POST, PUT, DELETE, etc.).
http.MethodPost is a predefined constant in Go's net/http package that represents the **POST HTTP method.**
The if statement checks if the incoming request's method is NOT a POST method.

Route Registration:

/shorten is the URL path/endpoint
When a request comes to this specific path, the createShortURL function will be called
This means any HTTP request to "/shorten" will be processed by the createShortURL function