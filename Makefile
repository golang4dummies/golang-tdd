default: run

.PHONY: default run

qrcoded:
	mkdir -p build
	GOPATH=`pwd` go build -o build/qrcoded src/qrcoded.go

build: qrcoded

run: qrcoded
	./build/qrcoded

format:
	go fmt src/$(package)/*.go

clean:
	rm -rf build
