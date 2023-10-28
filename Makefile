.PHONY: build
build:
	go build -o ./build/minigun ./cmd/minigun/*

run: build
	./build/minigun

.PHONY: vhs
vhs: build
	vhs ./demo/base.tape
