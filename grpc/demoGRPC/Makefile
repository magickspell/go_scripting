.PHONY: gen-go
gen-go:
	protoc -I=protos --go_out=go/pb protos/my/my.proto --go_opt=source_relative


# .PHONY: run-go
# run-go:
# 	cd ./go
# 	go run cmd/test/main.go