# Go gRPC Examples

Small Go examples for building gRPC clients and servers from Protocol Buffers.

## Examples

- `greet/` — unary and streaming RPCs, deadlines, and TLS.
- `calculator/` — a simple unary sum RPC.
- `blog/` — CRUD and server-streaming RPCs backed by MongoDB.

Each example has `proto/`, `server/`, and `client/` directories. Edit the `.proto`
file to define messages and RPCs; `*.pb.go` and `*_grpc.pb.go` are generated
files and should not be edited by hand. Each server embeds its generated
`pb.<Service>Server` interface, implements the RPC methods, and registers the
server with `pb.Register<Service>Server`. The client creates a connection and
uses the generated `pb.New<Service>Client`.

## Build and run

Install Go, `protoc`, and the Go protobuf plugins (`protoc-gen-go` and
`protoc-gen-go-grpc`). From the repository root:

```sh
make greet
./bin/greet/server
```

In another terminal, from the repository root:

```sh
./bin/greet/client
```

Run `make calculator` or `make blog` to generate and build those examples.
The blog server also requires MongoDB; start it with
`docker compose -f blog/docker-compose.yml up -d`.
`make` builds all examples, and `make test` builds them and runs the Go tests.
