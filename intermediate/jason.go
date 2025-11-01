package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName    string  `json:"name"`
	Age          int     `json:"age,omitempty"` // if no age assigned, age will be ommited
	EmailAddress string  `json:"email,omitempty"`
	Address      Address `json:"adress"`
	LastName     string  `json:"-"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	person := Person{
		FirstName: "Ahmed",
		Address:   Address{City: "New York", State: "NY"},
		LastName:  "Hazem", // this will never marshalled as we used - int he struct tag
	}
	fmt.Println(person)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

	person1 := Person{FirstName: "Hazem", Age: 25, EmailAddress: "ahazem@fakeemail.com",
		Address: Address{City: "Giza", State: "GZ"}}

	jsonData1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData1))
	jsonData3 := `{"full_name": "Jenny Doe", "emp_id": "0009", "age": 30, "address": {"city": "San Jose", "state": "CA"}}`

	var employeeFromJson Employee
	err = json.Unmarshal([]byte(jsonData3), &employeeFromJson) // if not using & it will make a duplicate and the modification won't be applied on the original struct
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println(employeeFromJson)
	fmt.Println("Jenny's Age increased by 5 years", employeeFromJson.Age+5)
	fmt.Println("Jenny's city:", employeeFromJson.Address.City)

	listOfCityState := []Address{
		{City: "New York", State: "NY"},
		{City: "San Jose", State: "CA"},
		{City: "Las Vegas", State: "NV"},
		{City: "Modesto", State: "CA"},
		{City: "Clearwater", State: "FL"},
	}

	fmt.Println(listOfCityState)
	jsonList, err := json.Marshal(listOfCityState)
	if err != nil {
		log.Fatalln("Error Marshalling to JSON:", err)
	}

	fmt.Println("JSON List:", string(jsonList))

	// Handling unknown JSON structures
	jsonData4 := `{"name": "John", "age": 30, "address": {"city": "New York", "state": "NY"}}`

	var data map[string]any

	err = json.Unmarshal([]byte(jsonData4), &data)
	if err != nil {
		log.Fatalln("Error Unmarshalling JSON:", err)
	}

	fmt.Println("Decoded/Unmarshalled JSON:", data)
	fmt.Println("Decoded/Unmarshalled JSON:", data["address"])
	fmt.Println("Decoded/Unmarshalled JSON:", data["name"])

}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}
