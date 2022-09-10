# BrainCore

The Social Robot core services not running on a robot

From https://grpc.io/docs/languages/go/quickstart/

https://grpc.io/docs/protoc-installation/

```bash
wget https://github.com/protocolbuffers/protobuf/releases/download/v21.5/protoc-21.5-linux-x86_64.zip
```

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative thesocialrobot/thesocialrobot.proto
```