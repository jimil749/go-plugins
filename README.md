# go-plugins

A comparision of [go plugin package](https://golang.org/pkg/plugin/) and other plugin implementations for Go.

Following packages have been used:

1. [hashicorp go-plugin](https://github.com/hashicorp/go-plugin) (via gRPC)
2. [pie plugin](https://github.com/natefinch/pie): (via rpc)
3. [Native Go plugin](https://golang.org/pkg/plugin/)

## Todo

Benchmarking

# Plugin

Each plugin framework has been used to build Key/Value store CLI where mechanism for storing and retrieving keys is pluggable.

1. Building hashicorp go-plugin

```
$ cd hashicorp-go-plugin

# Build the main CLI
$ go build -o kv

# build the grpc plugin
$ go build -o kv-go-grpc ./plugin


# This tells the KV binary to use the "kv-go-grpc" binary
$ export KV_PLUGIN="./kv-go-grpc"

# Read and write
$ ./kv put hello world

$ ./kv get hello
world
```

2. Building pie go plugin

```
$ cd pie-plugin

$ cd plugin_provider

# Build the plugin
$ go build -o plugin_provider

# Add the binary to $PATH or move the binary to master_provider
$ export PATH="path/to/plugin_provider:$PATH"

$ cd ../master_provider

# run the master
$ go run main.go
```

3. Native Go plugin
```
$ cd native-go-plugin

# Build the plugin
$ go build -buildmode=plugin -o kv.so

# Run main.go
$ go run ../main.go

```
