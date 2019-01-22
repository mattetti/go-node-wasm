# go -> wasm <- node

## Go Build

cd into src/wasm and `make build` will create a `lib.wasm` file that can then be loaded in node.js or a browser.

## Launch the node code

`GOOS=js GOARCH=wasm go run -exec=./src/node/go_js_wasm_exec ./src/wasm/main.go`
