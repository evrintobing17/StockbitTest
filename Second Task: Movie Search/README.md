# Movie Search

## About
This is microservice that able to handle two transports REST JSON HTTP and GRPC

This project will run rest and grpc on different port

HTTP
```Shell
localhost:8801
```
GRPC
```Shell
localhost:8802
```

## Compile
That you can start go applications via 
```Shell
$ go run cmd/main.go
```

## Example
This is request example for the HTTP
```Shell
$ curl http://localhost:8801/search\?title=sepeda

{"Search":[{"Title":"Sepeda Presiden","Year":"2021","imdbID":"tt15556636","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BOWUyMDUzNGYtYmRjMy00NDA4LTllMmYtMjMyZWRhMjZhZWIzXkEyXkFqcGdeQXVyNzY4NDQzNTg@._V1_SX300.jpg"}]}

```
