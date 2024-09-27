# protoc  -I ./ --go_out=plugins=grpc:./ --go_opt=paths=source_relative task.proto heartbeat.proto
protoc -I ./ --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative heartbeat.proto
