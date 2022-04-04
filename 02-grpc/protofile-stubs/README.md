# Tutorial

```
go mod init github.com/luankosaka1/{...}
```

```
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

```
go install
```

Execute after create proto/user.proto

```
protoc --proto_path=proto proto/*.proto --go_out=pb
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```

Execute server

```
go run cmd/services/user.go
```


## KTR0731 (https://github.com/ktr0731/evans)

Use to create server client

After installed server client

```
evans -r repl --host localhost --port {port server client | 50051}
```

## Issue message

```
protoc-gen-go: program not found or is not executable
--go_out: protoc-gen-go: Plugin failed with status code 1.
```

```
sudo apt install protobuf-compiler golang-goprotobuf-dev -y
export PATH="$PATH:$(go env GOPATH)/bin"
```