# Trajectory API Service gRPC protocol files

## 1. How to use.

### Golang

```sh
go get github.com/trajectoryjp/TrajectoryApiService@main
```

## 2. For developers

### prerequisite

- [buf](https://github.com/bufbuild/buf)  
- protoc([protobuf](https://github.com/protocolbuffers/protobuf))
- 各プログラミング言語でのprotocプラグイン
    - Goを利用する場合 protoc-gen-go, protoc-gen-go-grpc

Homebrewの場合、以下でインストールできます

```sh
brew install bufbuild/buf/buf protoc
# Goを利用する場合
brew install protoc-gen-go protoc-gen-go-grpc
```

### Instruction

1. ディレクトリ環境整備<br>
本リポジトリ(trajectoryApiService)を開発するサービスと同列のディレクトリに配置します。<br>以下のようなフォルダー構造になります。
```
<基準のフォルダ>
├── TrajectoryApiService
│    ├── SpatialIdentifier
│    │   └── Protocol
│    │       └── v1
│    ├── Trajectory
│    │   └── TrajectoryApiService
│    │       └── Protocol
│    │           └── v1
│    ├── build
│    │   └── github.com
│    │       └── trajectoryjp
│    │           └── TrajectoryApiService
│    │               ├── ...
│    └── cmd
└── <開発するサービス>
```

2. コンパイル
```sh
cd cmd
./build.sh
```

3. モジュール作成<br>
```sh
cd build/github.com/trajectoryjp/trjx_api_service
go mod init github.com/trajectoryjp/trjx_api_service
```

1. アプリケーション側の対応<br>
`go.mod`にreplaceを追加
```
replace github.com/trajectoryjp/trjx_api_service => ../trjx_api_service/build/github.com/trajectoryjp/trjx_api_service
```
