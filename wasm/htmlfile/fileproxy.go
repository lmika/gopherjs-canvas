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

