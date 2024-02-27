PHONY: all

grpc-python:
	python3 -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. ./proto/*.proto

grpc-go:
	protoc -I. --go_out=. --go-grpc_out=. ./proto/*.proto
