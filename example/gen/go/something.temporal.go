// Code generated by protoc-gen-temporal. DO NOT EDIT MANUALLY.
//
// Source: something.proto

package example

import (
	context "context"
	temporal "go.temporal.io/sdk/temporal"
	strconv "strconv"
	strings "strings"
)

const (
	temporalSDKVersion = temporal.SDKVersion
)

// Runtime validation of temporal SDK version. The overhead is minimal and only run
// one when the package is loaded. This is the only way we've found so far to
// properly ensure this generated code is compatible with the expected Temporal SDK
// version. We're currently trying to open a PR on the Go temporal SDK to add some
// simple constant simplifying code generation and hopefully we'll be able to
// switch to comptime assertion.
func init() {
	temporalVersionComponents := strings.Split(temporalSDKVersion, ".")
	major, err := strconv.Atoi(temporalVersionComponents[0])
	if err != nil {
		panic("failed to convert major SDK version to integer")
	}
	minor, err := strconv.Atoi(temporalVersionComponents[1])
	if err != nil {
		panic("failed to convert minor SDK version to integer")
	}
	patch, err := strconv.Atoi(temporalVersionComponents[2])
	if err != nil {
		panic("failed to convert patch SDK version to integer")
	}

	panicPrefix := "Temporal SDK (" + temporalSDKVersion + ")"

	if major < 1 || major > 1 {
		panic(panicPrefix + " major version is not supported expect minimum 1 major version")
	}
	if minor < 16 {
		panic(panicPrefix + " minor version is not supported expect minimum 16 minor version")
	}
	if patch < 0 {
		panic(panicPrefix + " patch version is not supported expect minimum 0 patch version")
	}
}

type Client interface {
	SendEmail(ctx context.Context, req *SendEmailRequest) error
	SendEmailWithRetention(ctx context.Context, req *SendEmailWithRetentionRequest) error
}

type client struct{}

func (c *client) SendEmail(ctx context.Context, req *SendEmailRequest) error {
	return nil
}

func (c *client) SendEmailWithRetention(ctx context.Context, req *SendEmailWithRetentionRequest) error {
	return nil
}

func NewEmailWorkerClient() Client {
	return &client{}
}
