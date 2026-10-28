# connect

Connect で API を提供するサーバーの proto を置く buf モジュールです。リポジトリ直下の既存の proto（grpc-gateway で REST を公開するもの）とは別のモジュールとして管理します。

## 構成

```text
connect/
└── trajectory/
    ├── uasl_designer/v1/   # UASL Designer
    └── <サービス>/v1/
```

* パッケージは `trajectory.<サービス>.v1`（小文字の snake_case）とし、このディレクトリからのパスと一致させる
* RPC のパスは `/trajectory.<サービス>.v1.<Service>/<Method>` になる。ディレクトリ名の `connect` はパッケージに含まれない
* API の設計は [AIP](https://google.aip.dev/) を参考にしている。標準メソッド（Get / Create / Update がリソースを返す、Delete が `google.protobuf.Empty` を返す）の形を lint で止めないよう、`RPC_REQUEST_RESPONSE_UNIQUE` と `RPC_RESPONSE_STANDARD_NAME` は除外している

## go_package などの言語別オプション

proto には `go_package` などの言語別オプションを書きません。利用する側が buf の managed mode で、自分のモジュール配下に生成します。

```yaml
# 利用側の buf.gen.yaml の例 (Go)
version: v2
managed:
  enabled: true
  disable:
    - module: buf.build/googleapis/googleapis
  override:
    - file_option: go_package_prefix
      value: example.com/yourapp/gen
inputs:
  - git_repo: https://github.com/trajectoryjp/trjx_api_service.git
    ref: <branch or commit>
    subdir: connect
    paths:
      - trajectory/uasl_designer
plugins:
  - local: protoc-gen-go
    out: gen
    opt: module=example.com/yourapp/gen
  - local: protoc-gen-connect-go
    out: gen
    opt: module=example.com/yourapp/gen
```

## 確認

buf の設定（`connect/buf.yaml`）と依存のロック（`connect/buf.lock`）はこのディレクトリにあり、リポジトリ直下の buf.yaml とは独立しています。コマンドはこのディレクトリで実行します（リポジトリのルートから実行すると、直下の buf.yaml が使われてしまう）。

```bash
cd connect
buf format --diff --exit-code
buf lint
buf breaking --against '../.git#branch=<比較するブランチ>,subdir=connect'
buf dep update   # 依存（googleapis）を更新するとき
```

`buf breaking` は、比較するブランチにも `connect/buf.yaml` がある場合にだけ使えます。
