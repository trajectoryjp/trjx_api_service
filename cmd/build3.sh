cd ../
#protoc --go_out=./build --go-grpc_out=./build *.proto
protoc --go_out=. --go-grpc_out=. trajectory/system_account_authentication/protocol/v1/*.proto
protoc --go_out=. --go-grpc_out=. uas/**/**/*.proto
protoc --go_out=. --go-grpc_out=. trajectory/gcs_manager/protocol/v1/*.proto
protoc --go_out=. --go-grpc_out=. trajectory/uss_manager/protocol/v1/*.proto
protoc --go_out=. --go-grpc_out=. trajectory/type/*.proto
