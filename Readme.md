Based on: https://www.youtube.com/watch?v=j-Zl42meo-4

gRPC stream server and stream client searching through github api

gRPC generating command:
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative repository/repository.proto
```