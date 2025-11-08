echo "Generating climate_data_service gRPC codes..."
protoc -I=./proto --go-grpc_out=./gen --go_out=./gen ./proto/*.proto

cd ./gen/github.com/hikata101/climate_data
go mod init github.com/hikata101/climate_data
go mod tidy