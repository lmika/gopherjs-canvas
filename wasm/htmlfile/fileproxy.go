package htmlfile

import "syscall/js"

type File struct {
	file js.Value
}

func FromInput(inputElement js.Value) *File {
	return &File{file: inputElement.Get("files").Index(0)}
}

func (fp *File) Name() string {
	return fp.file.Get("name").String()
}

func (fp *File) Size() int {
	return fp.file.Get("size").Int()
}

func (fp *File) WithBytes(bytesFn func(bytes []byte)) {
	var promiseResolve js.Func
	promiseResolve = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		defer promiseResolve.Release()

		println("Received arrayBuffer")

		arrayBuffer := args[0]
		uint8View := js.Global().Get("Uint8Array").New(arrayBuffer)

		println("Reading to byte slice")
		bytes := make([]byte, fp.Size())
		js.CopyBytesToGo(bytes, uint8View)

		println("Done")
		bytesFn(bytes)
		return js.Null()
	})

	arrayBufferPromise := fp.file.Call("arrayBuffer")
	println("Got arrayBuffer promise: ", arrayBufferPromise.String())

	r := arrayBufferPromise.Call("then", promiseResolve)
	println("Chaning promose: ", r.String(), " using ", promiseResolve.String())
}
