package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	past    = "daysago"
	forward = "daysfwd"
)

func main() {
	flag.Usage = usage
	z := flag.Bool("z", false, "time to zero")
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		result := time.Now()
		if *z == true {
			result = time.Date(result.Year(), result.Month(), result.Day(), 0, 0, 0, 0, time.Local)
		}
		printTime(result)
	} else if args[0] == "yesterday" {
		yesterday(*z)
	} else if args[0] == "tomorrow" {
		tomorrow(*z)
	} else if strings.HasSuffix(args[0], past) {
		nb, err := strconv.Atoi(strings.Replace(args[0], past, "", 1))
		if err != nil {
			usage()
		}
		countDays((nb * -1), *z)
	} else if strings.HasSuffix(args[0], forward) {
		nb, err := strconv.Atoi(strings.Replace(args[0], forward, "", 1))
		if err != nil {
			usage()
		}
		countDays(nb, *z)
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
	fmt.Println("Local: \t\t\t", date)
	fmt.Println("UTC: \t\t\t", date.UTC())
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
			usage()
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

var usageinfo string = `gotime is a timestamp tool for humans.

Usage:

	gotime [flags] [TIMESTAMP|EXPR]

flags:
	-z			Make the timestamp at 0 hour, 0 minute, 0 second of that day

TIMESTAMP:
	Timestamp can be in seconds, millioseconds, microseconds or nanoseconds

EXPR:
	Expressions can be shortcut words:
		tomorrow		timestamp 1 day in the future
		yesterday		timestamp 1 day in the past
		2daysago		timestamp 2 days in the past
		5daysfwd		timestamp 5 days in the future

`

func usage() {
	fmt.Println(usageinfo)
	os.Exit(2)
}
