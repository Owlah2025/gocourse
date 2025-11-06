package main

import (
	"crypto/rand"
	"crypto/sha256"

	//"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	password := "password123"

	// hash := sha256.Sum256([]byte(password))
	// fmt.Println(password)
	// fmt.Println(hash)

	// hash512 := sha512.Sum512([]byte(password))
	// fmt.Println(hash512)

	// fmt.Printf("SHA-256 Hash hex val: %x\n", hash)
	// fmt.Printf("SHA-512 Hash hex val: %x\n", hash512)
	salt, err := generateSalt()
	fmt.Println("Original Salt:", salt)
	fmt.Printf("Original Salt in hexadicimal: %x\n", salt)
	if err != nil {
		fmt.Println("Error generating salt:", err)
	}

	//Hash the password with salt
	signUpHash := hashPassword(password, salt)
	//Store the salt and hash in the database, Just printing it as now
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:", saltStr)    //simulate as storring in database
	fmt.Println("Hash:", signUpHash) //simulate as storring in database
	hashOriginalPassword := sha256.Sum256([]byte(password))
	fmt.Println("Hash of just the password string without salt:", base64.StdEncoding.EncodeToString(hashOriginalPassword[:]))
	//verify
	//retrieve the saltStr and decode it
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to decode salt:", err)
		return
	}
	loginHash := hashPassword("password123", decodedSalt)

	//Compare the stored signUpHash with the loginhash
	if signUpHash == loginHash {
		fmt.Println("Password is correct. You are logged in.")
	} else {
		fmt.Println("Login failed. Please check user credentials")
	}
}

//what is dictionary attacks and rainbow table attacks?!
//if two ussers made the same password their hashed values will be different due to different salts

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt) //takes exactly lens(salt) from rand.reader and puts in salt
	//in other words it makes a random number equal to the length of salt slice wich is 16
	if err != nil {
		return nil, err //returning slice[] and err
	}
	return salt, nil
}

// Function to hash password
func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...) //appending password to salt to make saltedPassword
	//NOTE:convert to byte slice then destructure because append accept variadic parameter
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
