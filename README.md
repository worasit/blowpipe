# Introduction

I am the go client implementation of [blowpipe.xyz](https://blowpipe.xyz).  Well, I will be.  Now I am the skeleton that doesn't really do anything.

## Installation

To install directly via `go get..`, Ensure `$GOPATH` and `$GOBIN` are setup correctly, then

```bash
go get github.com/getblowpipe/blowpipe
```

## Building

Alternatively you can build from source

```bash
git@github.com:getblowpipe/blowpipe.git blowpipe-go
cd blowpipe-go
make build
```

## Protobuf

Client <--> Server comms are gRPC. For this reason when any gRPC development work is carried out you must rebuild the generated code.  

Currently the .proto file is included in this repository; eventually it may be moved out to a separate module the client will depend on.  For now, if you're going to change the wire protocol, ensure you have the build tools:

Install the build tools

    make setup

Verify you can build

    make proto

## Build (everything including from .proto)

    make all
    
