## Runes and encoding in strings. 

In Go strings are just sequence of bytes. They can hold arbitary data. 

However, Go has a type rune, which is similar to int32, which means that rune is int32 which can hold any unicode code point. 


When you are working with strings and characters, you can use rune, because rune can break strings into individual characters. 

So basically, if "boots" string take 5 bytes of the data, one emoji can take upto 4 bbytes of length. 
