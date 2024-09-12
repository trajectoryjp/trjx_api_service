cd ../
#protoc --go_out=./build --go-grpc_out=./build *.proto
protoc --go_out=./build --go-grpc_out=./build trajectory/system_account_authentication/protocol/v1/*.proto
protoc --go_out=./build --go-grpc_out=./build uas/**/**/*.proto
protoc --go_out=./build --go-grpc_out=./build trajectory/type/*.proto
protoc --go_out=./build --go-grpc_out=./build trajectory/gcs_manager/protocol/v1/*.proto
protoc --go_out=./build --go-grpc_out=./build trajectory/uss_manager/protocol/v1/*.proto
protoc --go_out=./build --go-grpc_out=./build trajectory/trajectory_gcs_service/protocol/v1/*.proto

