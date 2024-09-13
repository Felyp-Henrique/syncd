all: clean test build

clean:
	rm -f syncd syncctl

test:
	for file_test in `find tests/ -name '*_test.go'`; do \
		go test $$file_test; \
	done;

test.verbose:
	for file_test in `find tests/ -name '*_test.go'`; do \
		go test -v $$file_test; \
	done;

build: \
	build.daemon build.client

build.daemon:
	go build -o syncd ./cmd/syncd/main.go

build.client:
	go build -o syncctl ./cmd/syncctl/main.go
