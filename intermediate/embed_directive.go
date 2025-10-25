package main

import (
	"embed" // using _ here is telling the compiler that it is ok , we use it just for its side effects and not using any of its functions
	"fmt"
	"io/fs"
	"log"
)

// this will embed example.txt in the final excutable
//
//go:embed example.txt
var content string

//go:embed embedFolder
var basicsFolder embed.FS

func main() {
	fmt.Println(content)
	fileContent, err := basicsFolder.ReadFile("embedFolder/embedFile.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	fmt.Println("file content: ", string(fileContent))

	//NOTE: if the walkdir finds a floder and a file it go through the folder first to print its content
	// then when backing from the folder it handels the files
	err = fs.WalkDir(basicsFolder, "embedFolder", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}
