package main

import (
	"fmt"
	"maps"
)

func main() {
	// var mapVariable map [keyType]valueType
	// mapVariable = make(map[keyTupe]valueType)

	// using a Map literal
	// mapVariable = map[keyType]valueType{
	// 	key1:value1,
	// 	key2:value2
	// }

	myMap := make(map[string]int)
	myMap["code"]= 1
	myMap["key1"]=2
	myMap["key2"]=3
	fmt.Println(myMap)
	
	delete(myMap,"key1")
	fmt.Println(myMap)

	// clear(myMap)
	// fmt.Println(myMap)

	// value,unknowValue := myMap["code"]
	// fmt.Println(value,unknowValue)

	_,ok := myMap["code"]
	if ok{
		fmt.Println(" there is a value associated to the 'code' key?")
	} else {
		fmt.Println("No value exists with 'code' key")
	}

myMap2 := map[string]int{"a":1, "b":2}
myMap3 := map[string]int{"a":1, "b":2}

if maps.Equal(myMap3,myMap2){
	fmt.Println("the two maps is equal")
}else{
	fmt.Println("the two maps is not equal")
}

for k,v:= range myMap3 {
	fmt.Println(k,v)
}

var nilMap  map[string]string 

if nil == nilMap {
	fmt.Println("the map is initialized with nil ")
}

k,v := nilMap["key"]
fmt.Println("is there is a value associated to 'key' key ?",v,"of the key:",k)

//very important part
//if i need to add key and values to the nilMap 
//nilMap["key"] = "value"  xxxxxxxxxxx wroooong -> this will not work
//we should use make() function 

nilMap= make(map[string]string )
nilMap["key"] = "value"
fmt.Println(nilMap["key"])


fmt.Println("nilMap length is: ",len(nilMap))
//the length of the map is the number of keys


map2D := make(map[string]map[string]string)
map2D["map1"] = nilMap
fmt.Println(map2D)

}