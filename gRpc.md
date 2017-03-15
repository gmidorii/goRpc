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

### What's rpc
`RPC = remote procedure call`   
ネットワーク越しの、サブルーチンを実行を可能にする技術。ネットワーク越しを意識することなく、
利用することができる。(通信等に関して記述しなくて良い)  
ローカルなサブルーチンと変わらずに、使うことができる。

### About grpc
```
gRPC is a modern open source high performance RPC framework 
that can run in any environment. 
```
- Rpcのフレームワーク
- どんな環境でも動く
- 様々な言語対応
- HTTP/2通信による双方向通信



### reference
- http://www.grpc.io/docs/quickstart/go.html
- https://ja.wikipedia.org/wiki/RPC
- http://www.grpc.io/about/