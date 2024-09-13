all: clean test build

clean:
	rm -f syncd syncctl

test:
	for test_files in `find tests/ -name '*_test.go'`; do \
		go test $$test_files; \
	done;

test.verbose:
	for test_files in `find tests/ -name '*_test.go'`; do \
		go test -v $$test_files; \
	done;

build: build.daemon build.client

build.daemon:
	go build -o syncd ./cmd/daemon/main.go

build.client:
	go build -o syncctl ./cmd/client/main.go
