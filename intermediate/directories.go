package main

import (
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// checkError(os.Mkdir("subdir", 0755))
	// defer os.RemoveAll("subdir")
	checkError(os.WriteFile("subfile.txt", []byte("hello "), 0755))
	checkError(os.MkdirAll("subdir/parent/child1", 0755))
	checkError(os.MkdirAll("subdir/parent/child2", 0755))
	checkError(os.MkdirAll("subdir/parent/child3", 0755))

}
