.PHONY: gen
gen:
	mkdir -p mocks
	go generate ./...


# 运行 ginkgo 或者 go test 均可运行测试
.PHONY: test
test:
	go test
