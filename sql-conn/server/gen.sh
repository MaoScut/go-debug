protoc -I . --grpc-gateway_out ./gen \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    service/service.proto
protoc -I . \
    --go_out ./gen --go_opt paths=source_relative \
    --go-grpc_out ./gen --go-grpc_opt paths=source_relative \
    service/service.proto