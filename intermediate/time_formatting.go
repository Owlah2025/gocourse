package main

import (
	"fmt"
	"time"
)

func main() {
	// reference time
	// Mon Jan 2 15:04:05 MST 2006

	layout := "2006-01-02T15:04:05z07:00"
	str := "2025-10-10T14:30:18z07:00"

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println(t)

	str1 := "Jul 03, 2024 03:18 PM" //2024-07-03 15:18:00 +0000 UTC time.time format
	layout1 := "Jan 02, 2006 03:04 PM"

	t1, err := time.Parse(layout1, str1)
	fmt.Println(t1)

}
