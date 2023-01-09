package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	then := time.Date(2009, 11, 17, 20, 35, 4, 0, time.UTC)

	fmt.Println(then)
	fmt.Println(then.Year(), then.Month(), then.Day(), then.Hour(), then.Minute(), then.Second(), then.Nanosecond(), then.Location())

	fmt.Println(then.Weekday())
	fmt.Println(then.Before(now))
	fmt.Println(then.After(now))

	diff := now.Sub(then)
	fmt.Println(diff)
	fmt.Println(diff.Hours(), diff.Minutes(), diff.Seconds())
	fmt.Println(then.Add(diff))
	fmt.Println(then.Add(-diff))
}
