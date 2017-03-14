# grpc

## introduction
1. `go get google.golang.org/grpc`
2. Install Protocol Buffers v3
3. environment variable PATH to include `protoc`

## build
```
 protoc -I rpc/ rpc/gorpc.proto --go_out=plugins=grpc:rpc
```

### Type(protocol buffers)
- enum
```prpto
enum Hoge {
    hoge1 = 0;
    hoge2 = 1;
}
```
- 


### reference
http://www.grpc.io/docs/quickstart/go.html
