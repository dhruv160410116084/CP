package main

import "fmt"

func main_events_conflict() {
	event1 := []string{"01:15", "02:00"}
	event2 := []string{"02:00", "03:00"}

	fmt.Println(event1[1] < event2[0])
	fmt.Println(event2[1] < event1[0])

}
