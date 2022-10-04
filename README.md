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

[![Project license](https://img.shields.io/github/license/lucasclerissepro/protoc-gen-temporal.svg?style=flat-square)](LICENSE)

[![Pull Requests welcome](https://img.shields.io/badge/PRs-welcome-ff69b4.svg?style=flat-square)](https://github.com/lucasclerissepro/protoc-gen-temporal/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22)
[![code with love by lucasclerissepro](https://img.shields.io/badge/%3C%2F%3E%20with%20%E2%99%A5%20by-lucasclerissepro-ff1414.svg?style=flat-square)](https://github.com/lucasclerissepro)

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
    };
    
    option (temporal.workflow).signals = {
      name: "cancel_something",
      request_type: "CancelSomethingRequest"
    };
    
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

> **[?]**
> If your work was funded by any organization or institution, acknowledge their support here.
> In addition, if your work relies on other software libraries, or was inspired by looking at other work, it is appropriate to acknowledge this intellectual debt too.
