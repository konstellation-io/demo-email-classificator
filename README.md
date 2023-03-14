# demo-email-classificator

## How to build your .krt file

``` sh
./scripts/build_krt.sh  
./scripts/build_krt.sh v2
```

## How to send manual requests to your deployed versions

Here are some examples:  

``` sh
./scripts/run_test.sh <k8s_namespace> <runtime_name> <version_name> <grpc_service> [<message_number>] [<message>] 
./scripts/run_test.sh kre demo classificator-v1 Classificator 1
./scripts/run_test.sh kre demo classificator-v1 Classificator 3
```
