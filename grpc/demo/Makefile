.PHONY: gen-py
gen-py:
	protoc -I=protos --python_out=./python/pb ./protos/my/my.proto

.PHONY: gen-go
gen-go:
	protoc -I=protos --go_out=go/pb protos/my/my.proto --go_opt=source_relative

.PHONY: gen
gen:
	buf generate protos

.PHONY: docker-gen
docker-gen:
	docker run --rm -v $(pwd):/workspace bufbuild/buf generate --template /workspace/buf.gen.yaml

.PHONY: run-py
run-py:
	python3 python/writeProto.py

.PHONY: run-go
run-go:
	cd ./go
	go run cmd/test/main.go