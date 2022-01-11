//go:build js
// +build js

package main

import (
	"github.com/lmika/image-canvas/wasm/imgprocess"
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
	select {}
}

//export loadImage
func loadImage() {
	imageElem := js.Global().Get("document").Call("querySelector", "#image-file")

	file := htmlfile.FromInput(imageElem)
	println("Loading the image: ", file.Name())
	println("Size: ", file.Size())

	println("Opening...")
	file.WithBytes(func(bts []byte) {
		img, err := imgprocess.ReadImage(bts)
		if err != nil {
			println(err)
		}

		r := img.Bounds()

		// Render
		w, h := r.Size().X, r.Size().Y
		println(w, " x ", h)

		renderedImage := make([]byte, w*h*4)
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				r, g, b, a := img.At(x, y).RGBA()
				renderedImage[(y*w+x)*4] = byte(r >> 8)
				renderedImage[(y*w+x)*4+1] = byte(g >> 8)
				renderedImage[(y*w+x)*4+2] = byte(b >> 8)
				renderedImage[(y*w+x)*4+3] = byte(a >> 8)
			}
		}

		// Render image
		canvas := js.Global().Get("document").Call("getElementById", "canvas")
		ctx := canvas.Call("getContext", "2d")

		jsData := js.Global().Get("Uint8Array").New(len(renderedImage))
		js.CopyBytesToJS(jsData, renderedImage)

		otherArray := js.Global().Get("Uint8ClampedArray").New(jsData)

		imgData := js.Global().Get("ImageData").New(otherArray, w, h)
		ctx.Call("putImageData", imgData, 20, 20)
	})
}
