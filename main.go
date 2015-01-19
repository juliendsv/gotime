package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {

	z := flag.Bool("z", false, "time to zero")

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		result := time.Now()
		if *z == true {
			result = time.Date(result.Year(), result.Month(), result.Day(), 0, 0, 0, 0, time.Local)
		}
		printTime(result)
	} else if args[0] == "help" {
		printUsage()
	} else if args[0] == "yesterday" {
		yesterday(*z)
	} else if args[0] == "tomorrow" {
		tomorrow(*z)
	} else if strings.HasSuffix(args[0], "daysago") {
		nb, err := strconv.Atoi(strings.Replace(args[0], "daysago", "", 1))
		if err != nil {
			printUsage()
		}
		countDays((nb * -1), *z)
	} else {
		showTime(args)
	}
}

func yesterday(z bool) {
	countDays(-1, z)
}
func tomorrow(z bool) {
	countDays(1, z)
}

func countDays(nb int, z bool) {
	result := time.Now().AddDate(0, 0, nb)
	if z == true {
		result = time.Date(result.Year(), result.Month(), result.Day(), 0, 0, 0, 0, time.Local)
	}

	printTime(result)
}

func printTime(date time.Time) {
	date = date.In(time.Local)
	fmt.Println("UTC: \t\t\t", date.UTC())
	fmt.Println("Local: \t\t\t", date)
	fmt.Println("timestamp: \t\t", date.Unix())
	fmt.Println("Milli timestamp: \t", date.Unix()*1000)
	fmt.Println("Nano timestamp: \t", date.UnixNano())
}

func showTime(args []string) {
	timestamp, _ := strconv.Atoi(args[0])
	if timestamp == 0 {
		t, err := time.Parse("2006-01-02 15:04:05", args[0])
		if err != nil {
			fmt.Printf("Error: unknown format \"%s\" \n", args[0])
			printUsage()
			return
		}
		printTime(t)
		return
	}
	date := guessTimestamp(int64(timestamp))
	printTime(date)
}

func guessTimestamp(ts int64) time.Time {
	switch {
	case ts <= math.MaxInt32:
		return time.Unix(ts, 0) // Seconds
	case ts <= math.MaxInt32*1e3:
		return time.Unix(0, ts*1e6) // Milliseconds
	case ts <= math.MaxInt32*1e6:
		return time.Unix(0, ts*1e3) // Microseconds
	default:
		return time.Unix(0, ts) // Nanoseconds
	}
}

func printUsage() {
	fmt.Println("Usage:\n")

	fmt.Println("gotime")
	fmt.Println("-----------------")
	printTime(time.Now())
	fmt.Println("\n")

	fmt.Println("gotime 1349034753")
	fmt.Println("-----------------")
	showTime([]string{"1349034753"})
	fmt.Println("\n")

	fmt.Println("gotime yesterday")
	fmt.Println("-----------------")
	yesterday(false)
	fmt.Println("\n")

	fmt.Println("gotime -z yesterday")
	fmt.Println("-----------------")
	yesterday(true)
	fmt.Println("\n")

	fmt.Println("gotime tomorrow")
	fmt.Println("-----------------")
	tomorrow(false)
	fmt.Println("\n")

	fmt.Println("gotime 5daysago")
	fmt.Println("-----------------")
	countDays(-5, false)
	fmt.Println("\n")

	fmt.Println("gotime 1405967017972502579")
	fmt.Println("-----------------")
	showTime([]string{"1405967017972502579"})
	fmt.Println("\n")

	fmt.Println("gotime \"2014-02-03 19:54:02\"")
	fmt.Println("-----------------")
	showTime([]string{"2014-02-03 19:54:02"})
	fmt.Println("\n")
}
