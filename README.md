# Raylib + Go Project
## Learning game development in go while tackling socket communication

In this project two main goals of mine where
 - to learn and work with go
 - make a simple tcp/ip server and client

This was achieved by creating a simple graphic using the raylib-go binding as well as networking using the `net` package in Go.


### Stuff Learned
 - Go functions, variables and control structures
 - Structs, object oriented programming style in go
 - Channels, Go routines, Mutex locks
 - Creating a tcp/ip server and client
 - Serializing data using gob
 - Building go programming and managing modules


### `net` Package

#### Server

A simple set of functions and constants for creating various network interfaces (TCP / IP, UDP). 
- `net.ListenTCP("tcp", addr)` to create new server
- returns listener pointer and error
- `listener.Accept` returns connection, err
- start go routine with connection as argument to concurrently handle each connection
- use `connection.Write(bytes)` to send data
- use `connection.Read(buffer)` to receive data
#### Client
- `net.Dial("tcp", addr)` returns connection and error
- use `conn.Write(bytes)` to send byte data
- read same as server

#### Go nuances

- := for type inferred assignement otherwise use var identifier datatype
- no while loop `for condition {}` for any loop
- datatypes are usually specified after identifiers
