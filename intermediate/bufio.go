package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// reader := bufio.NewReader(strings.NewReader("Hello, bufio packageee!\n"))
	// fmt.Printf("reader.Size(): %v\n", reader.Size())
	// {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	data, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println("Error reading from os.Stdin", err)
	// 	}
	// 	fmt.Println("data read:", data)
	// }
	// //reading byte slice
	// data := make([]byte, 20)
	// n, err := reader.Read(data)
	// if err != nil {
	// 	fmt.Println("Error reading:", err)
	// 	return
	// }
	// fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	// line, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error reading string:", err)
	// 	return
	// }
	// fmt.Println("Read string:", line)

	// Writer
	writer := bufio.NewWriter(os.Stdout)
	_ = writer

	//writing by slice
	data := []byte("Hello, bufio package! \n") // stored as:[72 101 108 108 ......] it made type conversion from string to int (ascii code)
	nn, err := writer.Write(data)
	if err != nil {
		fmt.Println("Error writing data:", err)
		return
	}
	fmt.Println("number of bytes of the data: ", nn)
	err = writer.Flush() //IMPORTANT: data is stored in an internal buffer
	//to improve performance by reducing the number of system calls.
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	// writing string
	str := "This is a string.\n"
	n, err := writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string", err)
		return
	}
	fmt.Printf("Wrote %d bytes.\n", n)
	//flush the buffer again
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	// fmt.Println(decimalToBinary(10))
	// binaryString, _ := fmt.Println(decimalToBinary(300))
	// fmt.Println("number of bits or runes in 300 in binary:", binaryString-1)
	// anyString, _ := fmt.Println("hello")
	// fmt.Println("number of runes in hello:", anyString-1)   // integer returned from println = #runes-1
}

// func decimalToBinary(decimalNumber int) string {
// 	binaryNumber := fmt.Sprintf("%b", decimalNumber)
// 	return binaryNumber
// }
