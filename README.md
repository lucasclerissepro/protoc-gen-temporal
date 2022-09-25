# Atome

Atome is a golang framework highly focused on code generation. It uses Protobuf for schema definition and enrich default Protobuf with tons of generators 
based on industry based practices.


## Temporal based framework

The whole framework is built on top of [Temporal](https://temporal.io) which is one of the best way to build reliable services nowadays.
Atome cares about how services communicates and provides opinionated way to implement routes. The inside of the service are entirely up to the developer, we do not provide any wrappers around databases or other 3rd party systems.
