package main

import (
	"fmt"
	"net/url"
)

func main() {

	// [sheme://][userinfo@]host[:port][/path][?query][#fragment]
	rawURL := "https://example.com:8080/path?query=param#fragment"

	parsedURL, err := url.Parse(rawURL)

	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println(parsedURL)
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("User:", parsedURL.User)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Port:", parsedURL.Port())
	fmt.Println("path:", parsedURL.Path)
	fmt.Println("Query:", parsedURL.Query()) //map
	fmt.Println("RawQuery:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)
	fmt.Println("RawFragment:", parsedURL.RawFragment)

	rawURL1 := "https://example.com/path?name=ahmed&age=25"
	parsedURL1, err := url.Parse(rawURL1)

	if err != nil {
		fmt.Println("Error parsing rawURL1:", err)
		return
	}
	//Query parameters
	queryParams := parsedURL1.Query()

	fmt.Println(queryParams)
	fmt.Println("Name:", queryParams.Get("name"), "\nAge:", queryParams.Get("age"))
	//NOTE: because it is a map you acces the value using a key "name" and "age"

	//Building URL
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}
	fmt.Println("before query:", baseURL)
	query := baseURL.Query()
	query.Set("name", "ahmed")
	baseURL.RawQuery = query.Encode()
	fmt.Println("after query:", baseURL)
	fmt.Println("Built URL:", baseURL.String())          // it will prints as string even if you didn't usin String() method
	fmt.Printf("without String() method:%#v\n", baseURL) //this prints it as struct

	//IMPORTANT: Another way to build url
	values := url.Values{} //this is a map adding key value
	values.Add("name", "ahmed hazem")
	values.Add("age", "30")
	values.Add("city", "Giza")
	values.Add("country", "Egypt")

	encodedQuery := values.Encode()
	fmt.Println(encodedQuery)

	baseURL1 := "https://example.com/search"
	fullURL := baseURL1 + "?" + encodedQuery
	fmt.Println(fullURL)
}
