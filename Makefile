create:
	# protoc --proto_path=proto proto/helloworld/*.proto --go_out=helloworld
	# protoc --proto_path=proto proto/helloworld/*.proto --go-grpc_out=helloworld
	protoc -I ./proto \
   --go_out ./proto --go_opt paths=source_relative \
   --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
   ./proto/helloworld/helloworld.proto

clean:
	rm proto/helloworld/*.go