package main

import (
	"os"
)

func main() {
	src, err := os.Open("27_code-in-process/27_package-os/01_Read/01/main.go")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = src.Close()
	}()

	dst, err := os.Create("27_code-in-process/27_package-os/01_Read/01/dst.txt")
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	info, _ := src.Stat()
	bs := make([]byte, info.Size())
	_, _ = src.Read(bs)
	_, _ = dst.Write(bs)
}

// this is a limit reader
// we limit what is read
// see io.ReadFull for func similiar to (*File)Read
