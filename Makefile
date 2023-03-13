#!/bin/bash
run: 
	@go run main.go

gen-proto:
	@protoc --proto_path=common common/model/*.proto --go_out=common/model --go-grpc_out=common/model