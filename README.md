# finpay-models
Finpay module for handling models

protoc user/proto/user.proto \
  --go_out=. \
  --go-grpc_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative
