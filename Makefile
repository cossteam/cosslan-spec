# 防止命令行参数被误认为是目标
%:
	@:

.PHONY: dep
dep: ## Get the dependencies
	@go mod tidy

.PHONY: lint
lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}

.PHONY: vet
vet: ## Run go vet
	go vet ./...

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: test
test: fmt vet## Run unittests
	@go test -short ./...

.PHONY: gen-grpc
gen-grpc:
	@for dir in $$(find api -mindepth 1 -maxdepth 1 -type d -not -path '*/\.*' -exec basename {} \;); do \
		echo "Generating grpc for $$dir"; \
		if [ -d api/$$dir/grpc/v1 ]; then \
			protoc -I=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false api/$$dir/grpc/v1/*.proto; \
			protoc-go-inject-tag -input=api/$$dir/grpc/v1/*.pb.go; \
		fi; \
	done
	go fmt ./...



