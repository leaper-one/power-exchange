## Generate rpc code
```sh
goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```
## Generate dockerfile
```sh
goctl docker -go user.go
```