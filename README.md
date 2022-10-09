<p align="center">
   <a href="https://dyn.gg">
   <img src="./img/banner.svg">
   </a>
</p>

<div align="center">
  Typesafe temporal workflows
  <br />
  <a href="#about"><strong>Explore the docs »</strong></a>
  <br />
  <br />
  <a href="https://github.com/lucasclerissepro/protoc-gen-temporal/issues/new?assignees=&labels=bug&template=01_BUG_REPORT.md&title=bug%3A+">Report a Bug</a>
  ·
  <a href="https://github.com/lucasclerissepro/protoc-gen-temporal/issues/new?assignees=&labels=enhancement&template=02_FEATURE_REQUEST.md&title=feat%3A+">Request a Feature</a>
  ·
  <a href="https://github.com/lucasclerissepro/protoc-gen-temporal/discussions">Ask a Question</a>
</div>

<div align="center">
<br />

<p align="center">

  <a aria-label="Dyn.gg logo" href="https://dyn.gg">
    <img src="https://img.shields.io/badge/MADE%20BY%20Dyn-000000.svg?style=for-the-badge">
  </a>

  <a aria-label="License" href="https://github.com/lucasclerissepro/protoc-gen-temporal/blob/main/license.md">
    <img alt="" src="https://img.shields.io/npm/l/next.svg?style=for-the-badge&labelColor=000000&color=F9B818">
  </a>

  <a aria-label="Join the community on GitHub" href="https://github.com/orgs/lucasclerissepro/protoc-gen-temporal/discussions">
    <img alt="" src="https://img.shields.io/badge/Join%20the%20community-orange.svg?style=for-the-badge&labelColor=000000&color=F9B818&logo=github">
  </a>
</p>

</div>

<h2><img height="20" src="./img/about.svg">&nbsp;&nbsp;About the project</h2>

> A protobuf plugin to automatically scaffold a temporal worker with sane
> defaults. On top of that this plugin can generate typesafe clients for
> consumers to call your workflow. 

<h2><img height="20" src="./img/installation.svg">&nbsp;&nbsp;Installation</h2>

You need to install [protobuf](https://developers.google.com/protocol-buffers)
to compile your protobuf. The easiest option to work with this plugin is to use
[buf](https://buf.build/).

<h2><img height="20" src="./img/gettingstarted.svg">&nbsp;&nbsp;Getting started</h2>

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

<h2><img height="20" src="./img/support.svg">&nbsp;&nbsp;Support</h2>

Reach out to the maintainer at one of the following places:

- [GitHub Discussions](https://github.com/lucasclerissepro/protoc-gen-temporal/discussions)
- Contact options listed on [this GitHub profile](https://github.com/lucasclerissepro)

<h2><img height="20" src="./img/contributing.svg">&nbsp;&nbsp;Contributing</h2>

First off, thanks for taking the time to contribute! Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make will benefit everybody else and are **greatly appreciated**.

Please read [our contribution guidelines](docs/CONTRIBUTING.md), and thank you for being involved!

<h2><img height="20" src="./img/security.svg">&nbsp;&nbsp;Security</h2>

protoc-gen-temporal follows good practices of security, but 100% security cannot be assured.
protoc-gen-temporal is provided **"as is"** without any **warranty**. Use at your own risk.

The following components are scanned:

- protoc plugin
- generated code

_For more information and to report security issues, please refer to our [security documentation](docs/SECURITY.md)._

<h2><img height="20" src="./img/license.svg">&nbsp;&nbsp;License</h2>

This project is licensed under the **MIT license**.

See [LICENSE](LICENSE) for more information.

<h2><img height="20" src="./img/ack.svg">&nbsp;&nbsp;Acknowledgements</h2>

- [Datadog](https://datadog.com): For inspiring this project.
