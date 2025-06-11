module github.com/tlsh0/grpc-todo-list/apigateway-service

go 1.24.2

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.3
	github.com/tlsh0/grpc-todo-list/task-service v0.0.0-20250611075838-c61aeab19295
	github.com/tlsh0/grpc-todo-list/user-service v0.0.0-20250611075838-c61aeab19295
	google.golang.org/genproto/googleapis/api v0.0.0-20250603155806-513f23925822
	google.golang.org/grpc v1.73.0
	google.golang.org/protobuf v1.36.6
)

require (
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250528174236-200df99c418a // indirect
)
