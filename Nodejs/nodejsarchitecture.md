# Nodejs arcitecture... 

Nodejs follows event-driven, non-blocking, I/O model which is based on V8 engine (Google Chrome)

# How are request being handled in nodejs server. 
There is a thing called **event queue** in which there are request being stored. 
Then that request will go to **event loop**

Work of event loop is to watch over the event queue and get the request in and out of that queue. First in, first out principle. 

## Blocking Operations and Non-Blocking Operations

#### Blocking -> Sync 
In blocking case, the request will go to the thread pool... (threads are like workers)
#### Non-Blocking -> ASync 
Event loop will check weather the particular request is blocking opreation or non-blocking operation.  

## V8 Engine (Execution Engine)
It's written in C++ and it converts js file to machine code by using JIT(Just in time) compilation. 

    It uses 2 main compilers internally. 
        1. Ignition :- This will generate byte-code.
        2. Turbofan :- This will convert bytecode into optimised machine code. 

## Nodejs binding layer. 
It is genrally recommended to use non-blocking requests in callback queues. 
