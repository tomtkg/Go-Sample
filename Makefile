TAG:=$(shell git rev-parse HEAD)
TIME_FILE:=test/time.json

.PHONY: build-separate-json
build-separate-json:
	go build -ldflags '-s -w' -o scripts/separate-json.exe cmd/separate-json/main.go

.PHONY: test-time-checker
test-time-checker:
	go test -v ./cmd/time-checker;
	go run cmd/time-checker/main.go < ${TIME_FILE};
	go run cmd/time-checker/main.go ${TIME_FILE};

.PHONY: docker-time-checker
docker-time-checker: docker-build-time-checker docker-run-time-checker

.PHONY: docker-build-time-checker
docker-build-time-checker:
	docker build -f docker/Dockerfile.time-checker -t time-checker:${TAG} .

.PHONY: docker-run-time-checker
docker-run-time-checker:
	docker run --rm -i time-checker:${TAG} < ${TIME_FILE}
