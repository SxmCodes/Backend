package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func httpserver (w http.ResponseWriter, r *http.Request){
	fmt.Println("Hello world!")
}
// we need to make an endpoint, specific endpoint set kar sakte hai, /generate-pdf jana client data send karega (curl ya html data)
func Takingnput(w http.ResponseWriter, r *http.Request){
	// Getting a post request. We will check weather the request is post request or not. 
	fmt.Println("Taking input")
	if r.Method != http.MethodPost{
		fmt.Println("Error")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body) // yeh pura body bytes mein dega. Bytes ko string mein convert kar sakta hai. It will returrors. 

	if err != nil {
		fmt.Println("Error occured!")
	}

	// json and html ko alag karna. converting that data into struct. 
	type JsonData struct {
		// json data.
		// in order to change the data from json to strings, we need to break it into parts, title ya content. 
		Title string `json:"title"`
		Content string `json:"content"`
	}
	var jsonData JsonData  
	err = json.Unmarshal(body, &jsonData)
	// checking error
	if err != nil {
		fmt.Println("Error occured!", err)
		http.Error(w, "Error in json", http.StatusBadRequest)
		return
	}

	// prinitng the data. 
	fmt.Println("Title: ", jsonData.Title)
	fmt.Println("Title: ", jsonData.Content)

	w.Write([]byte("Data received successfully!")) // giving response back to client.
	} // if there is html then no need to convert, we can use it straight way. 


func main(){
	http.HandleFunc( "/", httpserver)
	// This is to take input. 
	http.HandleFunc("/input", Takingnput)
	http.ListenAndServe(":8080", nil)
}