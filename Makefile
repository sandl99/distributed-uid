gen:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative gapi/*.proto

clean:
	rm gapi/*.pb.go

server:
	go run main.go --port=8001 --node=1 & go run main.go --port=8002 --node=2
