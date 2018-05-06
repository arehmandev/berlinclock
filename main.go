package main

import (
	"fmt"
	"os"
	"time"
)

// The int in each row represents the amount of lights on
type berlinClock struct {
	SecondBulb    bool
	FiveHourRow   int
	OneHourRow    int
	FiveMinuteRow int
	OneMinuteRow  int
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a time in the following format (24 hour):")
		fmt.Println(os.Args[0], "04:33:02 - Hour:Minute:Second")
		os.Exit(1)
	}
	// This is the 24 hour value we are trying to parse
	inputTime := os.Args[1]

	myclock := createBerlinClock(inputTime)

	printClock(*myclock)
}

func createBerlinClock(inputTime string) *berlinClock {

	// Put it into the correct format
	value := "Monday, 02-Jan-06 " + inputTime + " BST"

	// The form must conform to this format to be parsed
	form := "Monday, 02-Jan-06 15:04:05 BST"

	// Parse the string according to the form.
	parsedTime, err := time.Parse(form, value)
	if err != nil {
		fmt.Println("Error parsing time - invalid time given")
		fmt.Println("Please provide a time in the following format (24 hour):")
		fmt.Println(os.Args[0], "04:33:02 - Hour:Minute:Second")
		os.Exit(1)
	}

	myclock := new(berlinClock)

	// Get the Hours
	myclock.FiveHourRow = parsedTime.Hour() / 5
	myclock.OneHourRow = parsedTime.Hour() % 5

	// Get the minutes
	myclock.FiveMinuteRow = parsedTime.Hour() / 5
	myclock.OneMinuteRow = parsedTime.Hour() % 5

	// Get the seconds
	if parsedTime.Second()%2 == 0 {
		myclock.SecondBulb = true
	}

	return myclock
}

func printClock(myClock berlinClock) {

	fmt.Println("Seconds bulb on:", myClock.SecondBulb)
	fmt.Println("[ROW 1] Five hour row count:", myClock.FiveHourRow)
	fmt.Println("[ROW 2] One hour row count:", myClock.OneHourRow)
	fmt.Println("[ROW 3] Five minutes row count:", myClock.FiveMinuteRow)
	fmt.Println("[ROW 4] One minute row count:", myClock.OneMinuteRow)
}
