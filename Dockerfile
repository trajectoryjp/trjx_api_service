ARG VARIANT=1.24-bookworm
FROM golang:${VARIANT}

RUN go install github.com/bufbuild/buf/cmd/buf@v1.50.1 && \
    go install github.com/fullstorydev/grpcurl/cmd/grpcurl@v1.9.3 && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.5 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

COPY buf.* /proto_root/

WORKDIR /proto_root/github.com/trajectoryjp/trjx_api_service
