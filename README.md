# gRPC

First implementation of a gRPC Backend and Client.

## gRPC Server
Providing really simple APIs which will be accessible from the gRPC Client.
The recieved string will be written to the CockroackDB, that lives in the Cloud.

Currently, the ConnectionString is almost hard-written in the server. Only the Password will be passed from the env var.

## gRPC Client
Invoking Calls to the Server and passing a String 'Name' which will be handle from the gRPC Server.

## WebAssembly
In order to run the gRPC Client "stuff" from within the Browser, I want to have it run inside a WASM container.

## Goals
Once I managed the WASM-part, the ToDo-List will grow. 
Services and messages are getting defined inside the .proto-file, and the methods in the servers main.go file.


---

# When the App is running and technically working

## Tests
[] Unit Tests
[] Running Tests with Github Actions

## Build
[] Building a container for server, webassembly

## Release 
[] Release them to a container registry (ghcr.io)

## Deployment
[] Prepare for Kubernetes Deployment
[] ArgoCD Implementation (Apps-of-Apps)