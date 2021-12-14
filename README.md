```
mkdir grpc_example && cd $_
go mod init grpc_example

# protoを書く

protoc \
-I proto \
--go_out=./genproto \
--go_opt=paths=source_relative \
--go-grpc_out=./genproto \
--go-grpc_opt=paths=source_relative \
proto/*.proto

go get -u github.com/davecgh/go-spew/spew
go mod tidy

# serverを書く
# clientを書く

# サーバーの起動
go run server/main.go server/server.go

# クライアントからgRPCの呼び出し
go run client/unary/main.go
go run client/streaming/main.go

# grpcurlを使った確認
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
grpcurl -plaintext localhost:50051 list grpc_example.Hello
grpcurl -plaintext localhost:50051 grpc_example.Hello/Greet
grpcurl -plaintext -d '{"msg": "test"}' localhost:50051 grpc_example.Hello/StreamExample
```