all: clean build.daemon build.client

clean:
	rm -f syncd syncctl

build.daemon:
	go build -o syncd ./cmd/daemon/main.go

build.client:
	go build -o syncctl ./cmd/client/main.go