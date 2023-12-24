Generate protos using command below whilst in the /proto directory
```
protoc --proto_path=. --go_out=. --go-grpc_out=. transaction_*.proto
```
