package main

import (
	"fmt"
	"time"
)

func main() {

	tBegin := time.Now()
	//current local time
	fmt.Println(time.Now())

	//specific time
	specificTime := time.Date(2025, time.October, 24, 8, 30, 5, 87498853, time.UTC)
	fmt.Println(specificTime)

	//Parse time
	parsedTime, _ := time.Parse("06-01-02", "25-10-24") // NOTE: you can onlny this refernce Mon Jan 2 15:04:05 MST 2006
	//2006-01-02 15-04-05
	fmt.Println(parsedTime)

	t := time.Now()
	fmt.Println("Formatted time:", t.Format("2006-01-02 15-04-05")) //format works on time type and returns stirng type

	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println(oneDayLater)
	fmt.Println(time.Second.String())    //output: 1s
	fmt.Println(time.Second * 23)        //output: 23s
	fmt.Println(t.Day(), " ", t.Month()) //output: 17 october because today is 17 october 😀
	fmt.Println(t.Weekday())             //output Friday ! really? Woow 😒

	loc, _ := time.LoadLocation("Africa/Casablanca")

	tInMorocco := time.Now().In(loc)
	fmt.Println(loc)
	fmt.Println(tInMorocco)

	t1 := time.Date(2025, time.October, 17, 11, 0, 0, 0, time.UTC)
	t2 := time.Date(2025, time.October, 17, 11, 13, 25, 0, time.UTC)
	duration := t2.Sub(t1)

	fmt.Println("Duration", duration)

	tEnd := time.Now()
	fmt.Println("the duration of the program is:", tEnd.Sub(tBegin).Seconds())

	// compareTimes
	fmt.Println("t2 is after t2? ", t2.After(t1))
}
