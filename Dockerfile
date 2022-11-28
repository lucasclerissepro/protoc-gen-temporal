# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.19-alpine as builder

ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod go.sum ./

# download Go modules and dependencies
RUN go mod download

# copy directory files i.e all files ending with .go
COPY *.go ./

# compile application
RUN go build -o /protoc-gen-temporal
 
FROM scratch

LABEL "build.buf.plugins.runtime_library_versions.1.name"="google.golang.org/protobuf"
LABEL "build.buf.plugins.runtime_library_versions.1.version"="v1.27.1"

COPY --from=builder /protoc-gen-temporal /protoc-gen-temporal

ENTRYPOINT ["/protoc-gen-temporal"]
