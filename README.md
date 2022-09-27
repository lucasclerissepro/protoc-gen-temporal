# protoc-gen-temporal

A protobuf plugin simplifying work with [Temporal](https://temporal.io) by turning RPC Services into type safe workflows.

# Getting started

Starts by defining the following RPC service:

```
syntax = "proto3";

message DoSomethingRequest {
   hello string = 0;
}

service MyWorkflow {
  option (temporal.worker_defaults).task_queue = "my_workflow";
  
  // DoSomething is a workflow that does something
  rpc DoSomething(DoSomethingRequest) returns (google.protobuf.Empty) {
    option (temporal.workflow).queries = {
      name: "something_info",
      response_type: "SomethingInfoResponse"
    }
    
    option (temporal.workflow).signals = {
      name: "cancel_something",
      request_type: "CancelSomethingRequest"
    }
  }
}
```

It will generate a typesafe client SDK and well as interface that needs to be implemented on worker side.

# Language support

| Language   | Supported        |
|------------|------------------|
| Golang     | Work in progress |
| C#         | Not started      |
| Javascript | Not started      |
| Python     | Not started      |
