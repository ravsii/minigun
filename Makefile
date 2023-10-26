.PHONY: build
build:
	go build -o ./build/minigun .

run: build
	./build/minigun

.PHONY: vhs
vhs: build
	vhs ./demo/base.tape
