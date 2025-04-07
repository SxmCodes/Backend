package main

import (
	"fmt"
	"io"
	"net/http"
)

// creating proxyhandler
/*
	It takes response from the target server. 
	gives it back to the client with status code and response. 
	cleans my closing the response body!
*/
func proxyHandler(w http.ResponseWriter, r *http.Request){
client := &http.Client{}

// request
req, err := http.NewRequest(r.Method, "http://target-server.com" + r.URL.Path ,r.Body)
resp, err := client.Do(req) // this is how we make http request, the response body is the stream. If you don't close it you will get into leaks. 

if err!=nil {
	http.Error(w, "server error", http.StatusBadGateway)
	return

}
defer resp.Body.Close() // this is the response body from the target server. We ensures that this is closed but doesn't happen immediately. Defer will execute it at the very last of the function.
w.WriteHeader(resp.StatusCode) // the proxy sends back to the client. w is the respose, with the status code of the target server. 

io.Copy(w, resp.Body) // this copies the response to w, which we will send to user. 
 // this is the actual response data. 

}
// circuit breaker (if mehchanism is not working properly then request rok ko temp)
// works in 3 ways 1. closed, 2. open, 3. half-open

func main(){
}