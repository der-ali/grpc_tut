create: 
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

clean:
	rm proto/*.go