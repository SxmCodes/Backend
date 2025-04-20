SHA-256 algorithm works on raw binary data, and it expects bytes instead of strings. The hash.Hash interface's Write method expects a byte slice ([]byte) as input. That's why we need to convert the string s to a byte slice

Write is how we incopreate data into that hash function. It takes the byte slice and incoperate that byte into hash state. 
Memory efficient and Flexibility. 