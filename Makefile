all: clean test build

clean:
	sudo rm -rf syncd
	sudo rm -rf syncctl
	sudo rm -rf web/.next
	sudo rm -rf web/dist
	sudo rm -rf src/presentations/views/*.html
	sudo rm -rf src/presentations/views/*.ico
	sudo rm -rf src/presentations/views/*.txt
	sudo rm -rf src/presentations/views/_next

generate.mocks:
	go install github.com/golang/mock/mockgen

	for file_test in `find tests/ -name '*_test.go'`; do \
		go generate -v $$file_test; \
	done;

test: generate.mocks
	for file_test in `find tests/ -name '*_test.go'`; do \
		go test $$file_test; \
	done;

test.verbose: generate.mocks
	for file_test in `find tests/ -name '*_test.go'`; do \
		go test -v $$file_test; \
	done;

build: \
	build.daemon build.client build.web

build.daemon:
	go build -o syncd ./cmd/syncd/main.go

build.client:
	go build -o syncctl ./cmd/syncctl/main.go

build.web:
	cd web && npm run build
	sudo chmod -R 777 web/dist
	mv web/dist/* src/presentations/views/

execute.web:
	cd web && npm run dev
