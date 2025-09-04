#!/bin/zsh
if [[ ${PWD##*/} == "cmd" ]];then
    cd ../
fi

cd ../../../

if [ -d "github.com" ]; then
else
    echo "target directory('github.com') does not exist"
    exit(1)
fi

if ! command -v buf 2>&1 >/dev/null;then
    echo "using buf"
    buf dep update
    buf generate
else
    echo "using protoc"
    # installing dependency protoc-gen-openapiv2/options/annotations.proto via go cmd
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.27.2
    protoc \
        -I . \
        -I $(go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway/v2@v2.27.2) \
        --go_out=. --go-grpc_out=. \
        github.com/trajectoryjp/trjx_api_service/**/*.proto
fi

