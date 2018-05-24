package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// The int in each row represents the amount of lights on
type berlinClock struct {
	SecondBulb    int
	FiveHourRow   int
	OneHourRow    int
	FiveMinuteRow int
	OneMinuteRow  int
}

const (
	lampOn        = "Y"
	lampOff       = "O"
	fiveHourOn    = "RRRR"
	fiveHourOff   = "OOOO"
	oneHourOn     = "RRRR"
	oneHourOff    = "OOOO"
	fiveMinuteOn  = "YYRYYRYYRYY"
	fiveMinuteOff = "OOOOOOOOOOO"
	oneMinuteOn   = "YYYY"
	oneMinuteOff  = "OOOO"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a time in the following format (24 hour - Hour:Minute:Second ):")
		fmt.Println(os.Args[0], "04:33:02")
		os.Exit(1)
	}
	// This is the 24 hour value we are trying to parse
	twentyFourHour := os.Args[1]

	myclock := new(berlinClock)
	err := myclock.createBerlinClock(twentyFourHour)
	if err != nil {
		fmt.Println("Error parsing time - invalid time given")
		fmt.Println("Please provide a time in the following format (24 hour - Hour:Minute:Second ):")
		fmt.Println(os.Args[0], "04:33:02")
	}

	printClock(*myclock)

	fmt.Println(myclock.returnColor())
}

func (clock *berlinClock) createBerlinClock(inputTime string) (err error) {

	// Put it into the correct format
	value := "Monday, 02-Jan-06 " + inputTime + " BST"

	// The form must conform to this format to be parsed
	form := "Monday, 02-Jan-06 15:04:05 BST"

	// Parse the string according to the form.
	parsedTime, err := time.Parse(form, value)
	if err != nil {
		return err
	}

	// Get the Hours
	clock.FiveHourRow = parsedTime.Hour() / 5
	clock.OneHourRow = parsedTime.Hour() % 5

	// Get the minutes
	clock.FiveMinuteRow = parsedTime.Minute() / 5
	clock.OneMinuteRow = parsedTime.Minute() % 5

	// Get the seconds
	clock.SecondBulb = 1 - (parsedTime.Second() % 2)

	return nil
}

func printClock(myClock berlinClock) {

	fmt.Println("Seconds bulb on:", myClock.SecondBulb, "/1")
	fmt.Println("[ROW 1] Five hour row count:", myClock.FiveHourRow, "/ 4")
	fmt.Println("[ROW 2] One hour row count:", myClock.OneHourRow, "/ 4")
	fmt.Println("[ROW 3] Five minutes row count:", myClock.FiveMinuteRow, "/ 11")
	fmt.Println("[ROW 4] One minute row count:", myClock.OneMinuteRow, "/ 4")
}

func (clock *berlinClock) returnColor() string {

	a1 := returnSlice(lampOn)[:clock.SecondBulb]
	a2 := returnSlice(lampOff)[clock.SecondBulb:]

	b1 := returnSlice(fiveHourOn)[:clock.FiveHourRow]
	b2 := returnSlice(fiveHourOff)[clock.FiveHourRow:]

	c1 := returnSlice(oneHourOn)[:clock.OneHourRow]
	c2 := returnSlice(oneHourOff)[clock.OneHourRow:]

	d1 := returnSlice(fiveMinuteOn)[:clock.FiveMinuteRow]
	d2 := returnSlice(fiveMinuteOff)[clock.FiveMinuteRow:]

	e1 := returnSlice(oneMinuteOn)[:clock.OneMinuteRow]
	e2 := returnSlice(oneMinuteOff)[clock.OneMinuteRow:]

	allslices := [][]string{
		a1,
		a2,
		b1,
		b2,
		c1,
		c2,
		d1,
		d2,
		e1,
		e2,
	}

	return strings.Join(combineSlices(allslices), "")
}

func returnSlice(word string) []string {
	return strings.Split(word, "")
}

func combineSlices(twodslice [][]string) (returnslice []string) {

	for _, slice := range twodslice {
		for _, value := range slice {
			returnslice = append(returnslice, value)
		}
	}

	return
}
