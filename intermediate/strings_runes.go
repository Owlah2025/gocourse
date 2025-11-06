package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message :="Hello, \nGo!"
	message1 :="Hello, \tGo!"
	message2 :="Hello, \rGo!" //it returns the | to the first char then ovverwrite it 
	rowMessage := `Hello,\nGo!`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rowMessage)


	fmt.Println("the first letter in the message variable is:",message[0]) //Ascii
	fmt.Printf("the first letter in the message variable is %c \n",message[0])
	fmt.Println("the first letter in the message variable is:",string(message[0])) //casting to string


	greeting := "Hello "
	name := "Alice"
	fmt.Println(greeting+name)
	str1 := "Apple" //A has ASCII code of 65
	str2 := "apple" //a has ASCII code of 97
	str := "banana" //b has ASCII code of 98
	str3 := "app"		//a has ASCII code of 97

fmt.Println(str1<str2)
fmt.Println(str3<str1)
fmt.Println(str<str1)
fmt.Println(str<str3)

for i,char := range message {
	fmt.Printf("The charcter at index %d is %c\n",i,char)
	fmt.Println("message runes are: ",char)
}

fmt.Println("Runes count of greeting is:",utf8.RuneCountInString(greeting))
greetingWithName := greeting+name
fmt.Println(greetingWithName)

	var ch rune = 'a' // rune is an alias for int32
	jch := '&'
	fmt.Println(ch)
	fmt.Println(jch)

	sch := string (ch)
	fmt.Println("The typeof 'ch' is:",ch)
	fmt.Println("The typeof 'sch' is:",sch)

	fmt.Printf("the type of jch is %T\n",jch) // NOTE: the type is int32 because it is rune
	sjch := string(jch)
	fmt.Printf("the type of sjch is %T\n",sjch) // it is now string
}