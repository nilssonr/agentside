.PHONY: proto
proto:
	protoc -Iproto proto/authentication.proto \
		--go_opt=paths=source_relative \
		--go_out=api/pb \
		--go-grpc_out=api/pb \
		--go-grpc_opt=paths=source_relative