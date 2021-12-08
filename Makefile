.PHONY: separate-json
separate-json: build-separate-json run-separate-json

.PHONY: build-separate-json
build-separate-json:
	go build -o scripts/separate-json.exe cmd/separate-json/main.go

.PHONY: run-separate-json
run-separate-json:
	./scripts/separate-json.exe
