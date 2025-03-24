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
    buf dep update
    buf generate
else
    protoc --go_out=. --go-grpc_out=. github.com/trajectoryjp/trjx_api_service/**/*.proto
fi

