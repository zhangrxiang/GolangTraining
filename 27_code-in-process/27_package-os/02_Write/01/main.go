package main

import (
	"log"
	"os"
)

func main() {

	args := os.Args
	if len(args) <= 1 {
		args = append(args, "27_code-in-process/27_package-os/02_Write/01/test.txt")
	}
	dst, err := os.Create(args[1])
	if err != nil {
		log.Fatalf("error creating destination file:%v ", err)
	}
	defer dst.Close()
	dst.Write([]byte("Hello World"))
}

/*

os.Create
func Create(name string) (*File, error)

os
func (f *File) Write(b []byte) (n int, err error)

os
func (f *File) Read(b []byte) (n int, err error)

*/

/*

step 1 - at command line enter:
go install

step 2 - at command line enter:
programName initial.txt second.txt

*/
