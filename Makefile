.Phony: thewasm
thewasm:
	( cd wasm ; tinygo build -o ../public/js/main.wasm -target wasm ./main.go )
	# ( cd wasm ; GOOS=js GOARCH=wasm go build -o ../public/js/main.wasm ./main.go )

.Phony: run
run:
	go run .