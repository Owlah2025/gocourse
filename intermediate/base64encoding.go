package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("He~lo, Base64 Encoding")

	//Encode to Base 64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Encoded:", encoded)

	//Decode from Base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println("Decoded:", string(decoded))

	//URL safe, avoid '\' and '+'  if there is ~ in the string the encoded will contain +
	//so you can use the URLEncoding instead of StdEncoding

	urlSafeEncoding := base64.URLEncoding.EncodeToString(data)
	fmt.Println("Safe url encoding", urlSafeEncoding)
}

// usful in embedding small images or files directly into html or css using data urls
//NOTE: base64 encoding isn't encryption
