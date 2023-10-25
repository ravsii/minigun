build:
	go build -o ./build/minigun .

run: build
	./build/minigun

vhs: build
	vhs ./demo/base.tape
