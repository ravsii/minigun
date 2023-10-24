build:
	go build -o ./build/minigun .

run: build
	./build/minigun

vhs:
	vhs ./demo/base.tape
