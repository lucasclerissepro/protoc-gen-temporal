<div align="center">
  <h2>protoc-gen-temporal</h2>
  Typesafe temporal workflows
  <br />
  <a href="#about"><strong>Explore the docs »</strong></a>
  <br />
  <br />
  <a href="https://github.com/lucasclerissepro/protoc-gen-temporal/issues/new?assignees=&labels=bug&template=01_BUG_REPORT.md&title=bug%3A+">Report a Bug</a>
  ·
  <a href="https://github.com/lucasclerissepro/protoc-gen-temporal/issues/new?assignees=&labels=enhancement&template=02_FEATURE_REQUEST.md&title=feat%3A+">Request a Feature</a>
  .<a href="https://github.com/lucasclerissepro/protoc-gen-temporal/discussions">Ask a Question</a>
</div>

<div align="center">
<br />

<p align="center">

  <a aria-label="Dyn.gg logo" href="https://dyn.gg">
    <img src="https://img.shields.io/badge/MADE%20BY%20Dyn.gg-000000.svg?style=for-the-badge&logo=data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNTAiIGhlaWdodD0iNTAiIHZpZXdCb3g9IjAgMCA1MCA1MCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPGcgY2xpcC1wYXRoPSJ1cmwoI2NsaXAwXzQwMV8xNikiPgo8cGF0aCBkPSJNMTQuNDQ0NCAxNy43Nzc3SDBWMzIuMjIyMUgxNC40NDQ0VjE3Ljc3NzdaIiBmaWxsPSJ3aGl0ZSIvPgo8cGF0aCBkPSJNMzIuMjIyMSAwSDE3Ljc3NzdWMTQuNDQ0NEgzMi4yMjIxVjBaIiBmaWxsPSJ3aGl0ZSIvPgo8cGF0aCBkPSJNMzIuMjIyMSAxNy43Nzc3SDE3Ljc3NzdWMzIuMjIyMUgzMi4yMjIxVjE3Ljc3NzdaIiBmaWxsPSJ3aGl0ZSIvPgo8cGF0aCBkPSJNMzIuMjIyMSAzNS41NTU1SDE3Ljc3NzdWNTBIMzIuMjIyMVYzNS41NTU1WiIgZmlsbD0id2hpdGUiLz4KPHBhdGggZD0iTTQ1IDBIMzUuNTU1NVYxNC40NDQ0SDUwVjVDNTAgMi4yNDI4NyA0Ny43NTcxIDAgNDUgMFoiIGZpbGw9IndoaXRlIi8+CjxwYXRoIGQ9Ik01MCAxNy43Nzc3SDM1LjU1NTVWMzIuMjIyMUg1MFYxNy43Nzc3WiIgZmlsbD0id2hpdGUiLz4KPHBhdGggZD0iTTAgMzUuNTU1NVY0NUMwIDQ3Ljc1NzEgMi4yNDI4NyA1MCA1IDUwSDE0LjQ0NDRWMzUuNTU1NUgwWiIgZmlsbD0id2hpdGUiLz4KPHBhdGggZD0iTTM1LjU1NTUgMzUuNTU1NVY1MEg0NUM0Ny43NTcxIDUwIDUwIDQ3Ljc1NzEgNTAgNDVWMzUuNTU1NUgzNS41NTU1WiIgZmlsbD0id2hpdGUiLz4KPHBhdGggZD0iTTUgMEMyLjI0Mjg3IDAgMCAyLjI0Mjg3IDAgNVYxNC40NDQ0SDE0LjQ0NDRWMEg1WiIgZmlsbD0id2hpdGUiLz4KPC9nPgo8ZGVmcz4KPGNsaXBQYXRoIGlkPSJjbGlwMF80MDFfMTYiPgo8cmVjdCB3aWR0aD0iNTAiIGhlaWdodD0iNTAiIGZpbGw9IndoaXRlIi8+CjwvY2xpcFBhdGg+CjwvZGVmcz4KPC9zdmc+Cg==">
  </a>

  <a aria-label="License" href="https://github.com/lucasclerissepro/protoc-gen-temporal/blob/main/license.md">
    <img alt="" src="https://img.shields.io/npm/l/next.svg?style=for-the-badge&labelColor=000000">
  </a>
  <a aria-label="Join the community on GitHub" href="https://github.com/orgs/lucasclerissepro/protoc-gen-temporal/discussions">
    <img alt="" src="https://img.shields.io/badge/Join%20the%20community-blueviolet.svg?style=for-the-badge&labelColor=000000&logo=data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNTAiIGhlaWdodD0iNTAiIHZpZXdCb3g9IjAgMCA1MCA1MCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPGcgY2xpcC1wYXRoPSJ1cmwoI2NsaXAwXzQwMV8xNikiPgo8cGF0aCBkPSJNMTQuNDQ0NCAxNy43Nzc3SDBWMzIuMjIyMUgxNC40NDQ0VjE3Ljc3NzdaIiBmaWxsPSJ3aGl0ZSIvPgo8cGF0aCBkPSJNMzIuMjIyMSAwSDE3Ljc3NzdWMTQuNDQ0NEgzMi4yMjIxVjBaIiBmaWxsPSJ3aGl0ZSIvPgo8cGF0aCBkPSJNMzIuMjIyMSAxNy43Nzc3SDE3Ljc3NzdWMzIuMjIyMUgzMi4yMjIxVjE3Ljc3NzdaIiBmaWxsPSJ3aGl0ZSIvPgo8cGF0aCBkPSJNMzIuMjIyMSAzNS41NTU1SDE3Ljc3NzdWNTBIMzIuMjIyMVYzNS41NTU1WiIgZmlsbD0id2hpdGUiLz4KPHBhdGggZD0iTTQ1IDBIMzUuNTU1NVYxNC40NDQ0SDUwVjVDNTAgMi4yNDI4NyA0Ny43NTcxIDAgNDUgMFoiIGZpbGw9IndoaXRlIi8+CjxwYXRoIGQ9Ik01MCAxNy43Nzc3SDM1LjU1NTVWMzIuMjIyMUg1MFYxNy43Nzc3WiIgZmlsbD0id2hpdGUiLz4KPHBhdGggZD0iTTAgMzUuNTU1NVY0NUMwIDQ3Ljc1NzEgMi4yNDI4NyA1MCA1IDUwSDE0LjQ0NDRWMzUuNTU1NUgwWiIgZmlsbD0id2hpdGUiLz4KPHBhdGggZD0iTTM1LjU1NTUgMzUuNTU1NVY1MEg0NUM0Ny43NTcxIDUwIDUwIDQ3Ljc1NzEgNTAgNDVWMzUuNTU1NUgzNS41NTU1WiIgZmlsbD0id2hpdGUiLz4KPHBhdGggZD0iTTUgMEMyLjI0Mjg3IDAgMCAyLjI0Mjg3IDAgNVYxNC40NDQ0SDE0LjQ0NDRWMEg1WiIgZmlsbD0id2hpdGUiLz4KPC9nPgo8ZGVmcz4KPGNsaXBQYXRoIGlkPSJjbGlwMF80MDFfMTYiPgo8cmVjdCB3aWR0aD0iNTAiIGhlaWdodD0iNTAiIGZpbGw9IndoaXRlIi8+CjwvY2xpcFBhdGg+CjwvZGVmcz4KPC9zdmc+Cg==">
  </a>
</p>

</div>

## About

> A protobuf plugin to automatically scaffold a temporal worker with sane
> defaults. On top of that this plugin can generate typesafe clients for
> consumers to call your workflow. 

## Getting Started

### Prerequisites

You need to install [protobuf](https://developers.google.com/protocol-buffers)
to compile your protobuf. The easiest option to work with this plugin is to use
[buf](https://buf.build/).

## Usage

Starts by defining the following RPC service:

```protobuf
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

## Support

Reach out to the maintainer at one of the following places:

- [GitHub Discussions](https://github.com/lucasclerissepro/protoc-gen-temporal/discussions)
- Contact options listed on [this GitHub profile](https://github.com/lucasclerissepro)

## Contributing

First off, thanks for taking the time to contribute! Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make will benefit everybody else and are **greatly appreciated**.


Please read [our contribution guidelines](docs/CONTRIBUTING.md), and thank you for being involved!

## Authors & contributors

The original setup of this repository is by [Lucas Clerisse](https://github.com/lucasclerissepro).

For a full list of all authors and contributors, see [the contributors page](https://github.com/lucasclerissepro/protoc-gen-temporal/contributors).

## Security

protoc-gen-temporal follows good practices of security, but 100% security cannot be assured.
protoc-gen-temporal is provided **"as is"** without any **warranty**. Use at your own risk.

_For more information and to report security issues, please refer to our [security documentation](docs/SECURITY.md)._

## License

This project is licensed under the **MIT license**.

See [LICENSE](LICENSE) for more information.

## Acknowledgements

- [Datadog](https://datadog.com): For inspiring this project.
