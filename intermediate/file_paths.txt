package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	relativePath := "./data/file.txt"
	absolutePath := "/home/user/docs/file.txt"

	// Join paths using filepath.join
	joinedPath := filepath.Join("home", "Documents", "downloads", "file.zip")
	fmt.Println("Joined Path:", joinedPath)

	//clean paths to remove redundants
	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Normalized Path:", normalizedPath)

	dir, file := filepath.Split("/home/user/docs/file.txt")
	dir1, file1 := filepath.Split(absolutePath)
	fmt.Println("Dir1:", dir1, "& file1:", file1)
	fmt.Println("File:", file)
	fmt.Println("Dir:", dir)
	fmt.Println("Base of /home/user/docs/ is: ", filepath.Base("/home/user/docs/"))

	fmt.Println("Is relativePath variable absolute:", filepath.IsAbs(relativePath))
	fmt.Println("Is absolutePath variable absolute:", filepath.IsAbs(absolutePath))

	fmt.Println(filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	rel, err := filepath.Rel("a/b", "a/b/t/file") //retuns the target as relative path to the base
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Absolute Path:", absPath)
	}
}
