# Introduction

I am the go client implementation of [blowpipe.xyz](https://blowpipe.xyz).

## Installation

To install directly via `go get..`, Ensure \$GOPATH and \$GOBIN are setup correctly, then:

```bash
go get github.com/getblowpipe/blowpipe
```

## Building

Alternatively you can build from source

    git@github.com:getblowpipe/blowpipe.git
    cd blowpipe
    make build

## Protobuf

The Client <--> Server comms are gRPC.  For this reason when any gRPC development work is carried out you must rebuild the generated code. 

Install the build tools

    make setup

Verify you can build

    make proto

# Build

    make all
    
