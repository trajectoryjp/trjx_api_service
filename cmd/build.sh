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

# googleapis を準備し、そのパスを標準出力に返す関数
prepare_googleapis() {
  local version="master" # googleapis のバージョン (ブランチ名 or タグ名. commit hash不可)
  local build_dir="$(pwd)/.tmp"
  local googleapis_dir="${build_dir}/googleapis"

  if [ ! -d "${googleapis_dir}/.git" ]; then
    # ログ出力(echo)はstderrにリダイレクト
    echo ">>> Not a valid git repo. Re-cloning..." >&2
    rm -rf "$googleapis_dir"
    git clone https://github.com/googleapis/googleapis.git "$googleapis_dir" >/dev/null 2>&1
  else
    echo ">>> googleapis repository already exists. Fetching updates..." >&2
    (cd "$googleapis_dir" && git fetch) >/dev/null 2>&1
  fi

  echo ">>> Checking out specific commit: ${commit_hash}" >&2
  # gitコマンド自体の出力も捨てる(/dev/null)
  (cd "$googleapis_dir" && git checkout "$commit_hash") >/dev/null 2>&1

  echo "$googleapis_dir"
}

if ! command -v buf 2>&1 >/dev/null;then
    echo "using buf"
    buf dep update
    buf generate \
      && echo "build success"
else
    echo "using protoc"
    # installing dependency protoc-gen-openapiv2/options/annotations.proto via go cmd
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.27.2
    GOOGLEAPIS_DIR=$(prepare_googleapis)
    protoc \
        -I . \
        -I $(go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway/v2@v2.27.2) \
        -I "${GOOGLEAPIS_DIR}" \
        --go_out=. --go-grpc_out=. \
        github.com/trajectoryjp/trjx_api_service/**/*.proto \
      && echo "build success"
fi



