
TEMP_MAIN := temp_main.go
GEN_FILE := scripts/generate.go

generate-types:
	./scripts/fetch_api.sh
	@echo "package main" > $(TEMP_MAIN)
	@echo 'import "gateway-go-sdk/scripts"' >> $(TEMP_MAIN)
	@echo 'func main() { scripts.GenerateTypes() }' >> $(TEMP_MAIN)
	go run $(TEMP_MAIN)
	@rm -f $(TEMP_MAIN)
	

format:
	go fmt ./...

test:
	go test -count=1 -coverprofile coverage.out -coverpkg=./... ./...
	cat coverage.out | grep -v "scripts" | grep -v "othername" > coverage.text

.PHONY: generate-types format