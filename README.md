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
    <img src="https://img.shields.io/badge/MADE%20BY%20Dyn-000000.svg?style=for-the-badge&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAFEAAABUCAYAAAD6QtuFAAAACXBIWXMAAAsTAAALEwEAmpwYAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAANnSURBVHgB7dztbdswEAbg10UHyAjMBskGygRNJ6g6QbpB1AmSDeoN0k5gb5AR5A3qDd6SkNsCiU3q446UoHsAQz8ckCZNUUfyHMAYY4wxxhhjjDFG2QbKSF75yxXKOW42myMU5ejE1l8cytn7TryDog9Q5DuwRtkODKrT51CjNhL9B3f+skP5TgzC7XytdVtrjsQG8+jAIMzJ36BEZSSeRmGLeQmj8NaPxgOEfYSOJ8jbYrp7/3rG3IVJnPIarEkIaSirZTc9rINv7APl1ViLMFqoMAqxJr7BW8q7wVqwG4XSfmAhROJE3+CwMqkg6/pSTMdunnxEP4fYe76OryiNBUIays69FUpj5oeJf/+Rslp223VlKDQoqCP1acy9QYMSTg36TVmviTo1IoAgtCP/aFRqkIvUd0NdGuv9y1ggpKH83HtOhVwUGhRuJxepTyMCOGeHHJg5pKHOcjLmHpp8BVdcfkiT/Dwc+JAZejzwAPkt/++X3mB3izfIy0HrKIE6D5OXRJ1aIU1KdI6e0om5Q5p7lvUCSdSJ0eYQ0qRUkMLMW/7MF9Kk9Ap5kg8Wdt+Gg6x94ujyAfNQscdoTHaib+zeX/aQFea7WBjxC/NwOLV/uvBtUN5Tos6W5dV9+qf3zrYvcOsvX3r86fH06uP2Un4Mu9sozzLsPPlsMvbf+hILVH1ZO5bjoMEX3CQqFj3ipN4mbIreIRnTa+cawnyZz8yrpXbGBS+vJFS+PXZfnPQOeoxaCt7bhu3OVO6gJDSMeeTLuOD7kEc9Xc3X8Up9NUYYfXh/6riwsjj4113koD1scvZJBwlZ/s+R+irohjzb7Af5/D9X1ZG/GborXSfq3FGPQwk9Gj10+6xlZDlInWPaYJ55Pxwf4zWJclOx6lAt55pEyvG3XurkTzrkqTGR1q8Han+Zcov89JP850T5fbPCLnHodmmuMUeU2YGpsFaUO+IsuYNTDuWfnnmWYBNJ/yytguzPcj9hjSgbEDusEeWOEtaV+P4Wu4PvKYlBB8TX4xXGJdqHkGaLJeD0B0yTKLvlOA5LwvHLM60sscVMD/9wfBpeHSlz7HpcJ0EpBw5PB9klyhu6K/RXgyXjsJDHRcoZmyWmvuU/q/+LE8vPYfdjyTGB/EHjXxYYY4wxxhhjjDHGmKX5AwskxsWjl389AAAAAElFTkSuQmCC">
  </a>

  <a href="https://wakatime.com/badge/github/lucasclerissepro/protoc-gen-temporal">
    <img src="https://wakatime.com/badge/github/lucasclerissepro/protoc-gen-temporal.svg?style=for-the-badge&color=000000" alt="wakatime">
  </a>

  <a aria-label="License" href="https://github.com/lucasclerissepro/protoc-gen-temporal/blob/main/license.md">
    <img alt="" src="https://img.shields.io/npm/l/next.svg?style=for-the-badge&labelColor=000000&color=F9B818">
  </a>

  <a aria-label="Join the community on GitHub" href="https://github.com/orgs/lucasclerissepro/protoc-gen-temporal/discussions">
    <img alt="" src="https://img.shields.io/badge/Join%20the%20community-orange.svg?style=for-the-badge&labelColor=000000&color=F9B818&logo=github">
  </a>
</p>

</div>

<h2><img height="20" src="./img/whatissurreal.svg">&nbsp;&nbsp;About the project</h2>

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
    }
    
    option (temporal.workflow).signals = {
      name: "cancel_something",
      request_type: "CancelSomethingRequest"
    }
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

_For more information and to report security issues, please refer to our [security documentation](docs/SECURITY.md)._

<h2><img height="20" src="./img/license.svg">&nbsp;&nbsp;License</h2>

This project is licensed under the **MIT license**.

See [LICENSE](LICENSE) for more information.

## Acknowledgements

- [Datadog](https://datadog.com): For inspiring this project.
