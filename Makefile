TAG:=$(shell git rev-parse HEAD)
TIME_FILE:=test/time.json

.PHONY: separate-json
separate-json: build-separate-json run-separate-json

.PHONY: build-separate-json
build-separate-json:
	go build -o scripts/separate-json.exe cmd/separate-json/main.go

.PHONY: run-separate-json
run-separate-json:
	./scripts/separate-json.exe

.PHONY: time-checker
time-checker: build-time-checker run-time-checker

.PHONY: build-time-checker
build-time-checker:
	go build -o scripts/time-checker.exe cmd/time-checker/main.go

.PHONY: run-time-checker
run-time-checker:
	./scripts/time-checker.exe < ${TIME_FILE}

.PHONY: docker-time-checker
docker-time-checker: docker-build-time-checker docker-run-time-checker

.PHONY: docker-build-time-checker
docker-build-time-checker:
	docker build -f docker/Dockerfile.time-checker -t time-checker:${TAG} .

.PHONY: docker-run-time-checker
docker-run-time-checker:
	docker run --rm -i time-checker:${TAG} < ${TIME_FILE}
