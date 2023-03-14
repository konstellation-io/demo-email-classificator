# demo-email-classificator

## Requirements

* [pyhton](https://www.python.org/) >=v3.8
* [golang](https://go.dev/) >=v1.18 
* [protobuf-compiler](https://grpc.io/docs/protoc-installation/)
* [Go GRPC plugins](https://grpc.io/docs/languages/go/quickstart/)
* [grpcurl](https://github.com/fullstorydev/grpcurl)

## How to build your .krt file

```sh
./scripts/build_krt.sh  
./scripts/build_krt.sh v2
```

## Generate protobuf code

Execute the following script if changes occur in `.proto` files.

```sh
./scripts/generate_protobuf_code.sh
```

## How to send manual requests to your deployed versions

Here are some examples:  

```sh
./scripts/run_test.sh <k8s_namespace> <runtime_name> <version_name> <grpc_service> [<message_number>] [<message>] 
./scripts/run_test.sh kre demo classificator-v1 Classificator 1
./scripts/run_test.sh kre demo classificator-v2 Classificator 3
```
