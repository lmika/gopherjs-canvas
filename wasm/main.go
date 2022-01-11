//go:build js
// +build js

package main

import (
	"syscall/js"

	"github.com/lmika/image-canvas/wasm/htmlfile"
)

// This calls a JS function from Go.
func main() {
	println("adding two numbers.. eventually")

	res := js.Global().Get("document").Call("getElementById", "wasm-message")
	res.Set("innerText", "WASM loaded successfully")
}

//export loadImage
func loadImage() {
	imageElem := js.Global().Get("document").Call("querySelector", "#image-file")

	file := htmlfile.FromInput(imageElem)
	println("Loading the image: " + file.Name())
}
