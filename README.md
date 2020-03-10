# oriClient

## Overview

oriClient is a CLI that makes gRPC calls to the `oriServer`, a component part of the Ori tech test.

## Pre-requisites

1. `go` must be installed, version 1.13.x, to make use of `go modules`. For installation details, please see here:
https://golang.org/doc/install
2. `oriServer` must be running as a separate process, details of `oriserver` including installation and spin-up 
instructions can be found here: https://github.com/jon-wade/oriServer
 
## Install
 
 1. Clone the repo with `git clone https://github.com/jon-wade/oriClient.git`
 2. `cd oriClient` to change directories to the newly cloned repo
 3. `go mod download` to install the dependencies
 4. `go build` to create the binary
 5. `./oriClient -h` for instructions on the syntax of the CLI
 
```
oriClient % ./oriClient -h
Usage of ./oriClient:
  -host string
        hostname of oriserver, e.g. localhost (default "localhost")
  -method string
        math helper method, e.g. summation or factorial
  -port int
        port number of oriserver, e.g. 50051 (default 50051)
```

## Running tests

The repo contains E2E tests as well as Unit tests. To run the entire test suite, or the E2E tests `oriServer` must 
be running as a separate process on its default host and port. The unit tests can be run with no server in place.

* Run all tests `go test ./...`
* Run only E2E tests `go test -run E2E ./...`
* Run only Unit tests `go test -run Unit ./...`

## Tech test documentation

The documentation requirement of the tech test can be found within the `oriServer` repo `README.md` file, available here:
https://github.com/jon-wade/oriServer/blob/master/README.md