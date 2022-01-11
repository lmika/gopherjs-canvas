//go:build js
// +build js

package main

import (
	"strconv"
	"syscall/js"

	"github.com/lmika/image-canvas/wasm/htmlfile"
)

// This calls a JS function from Go.
func main() {
	println("adding two numbers.. eventually")

	res := js.Global().Get("document").Call("getElementById", "wasm-message")
	res.Set("innerText", "WASM loaded successfully")

	js.Global().Set("wasmExports", js.ValueOf(map[string]interface{}{
		"loadImage": js.FuncOf(func(this js.Value, args []js.Value) interface{} { loadImage(); return js.Null() }),
	}))

	// Prevent main from terminating.  This keeps the wasm program alive
	<-(make(chan struct{}))
}

//export loadImage
func loadImage() {
	imageElem := js.Global().Get("document").Call("querySelector", "#image-file")

	file := htmlfile.FromInput(imageElem)
	println("Loading the image: ", file.Name())
	println("Size: ", file.Size())

	println("Opening...")
	file.WithBytes(func(bts []byte) {
		println("Opened")
		println("Size = " + strconv.Itoa(len(bts)))
	})
}
