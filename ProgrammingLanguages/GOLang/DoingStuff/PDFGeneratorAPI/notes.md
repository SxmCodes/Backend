This methods accepts two arguments: the first is a []byte which represents the JSON object to unmarshal, and the second is any (introduced in Go 1.18 as an alias for interface{}) which should be a pointer to the target data structure for storing the result of unmarshalling the JSON data.

### Why are we using struct in place of maps?
Fixed Structure: Struct tab use hota hai jab tujhe pata hai JSON ka shape pehle se—jaise ```{ "title": "Test", "content": "Hello" }```. Tu chahta hai ki fields defined hon ```(title string, content string)```, taaki galat data aaye toh error mile.

Easy Access: Struct ke saath data.Title likhke direct access milta hai, code clean dikhta hai.

### Maps Kyun Use Kar Sakte Hain ?
Flexible Structure: Agar JSON ka shape fixed nahi hai—jaise kabhi ```{ "title": "Test" }, kabhi { "title": "Test", "age": 25, "color": "blue" }```—toh maps better hain. Map mein dynamically key-value store hota hai.

## Why we aren't using maps?
**Type Casting**: Map mein values **interface{} type ki hoti hain**, toh har baar check karna padta hai ki value string hai ya int—jaise value.(string) ya value.(int). Galat cast kiya toh runtime panic ho sakta hai.

Messy Code: ```data["title"]``` likhna padta hai, aur agar key galat likhi (jaise "titel"), toh runtime pe hi pata chalega—struct jaisa safety nahi.