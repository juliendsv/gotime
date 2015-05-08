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

var (
	past    = []string{"daysago", "dayago"}
	forward = []string{"daysfwd", "dayfwd"}
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
	} else if found, isit := stringHasSuffix(args[0], past); isit == true {
		nb, err := strconv.Atoi(strings.Replace(args[0], found, "", 1))
		if err != nil {
			usage()
		}
		countDays((nb * -1), *z)
	} else if found, isit := stringHasSuffix(args[0], forward); isit == true {
		nb, err := strconv.Atoi(strings.Replace(args[0], found, "", 1))
		if err != nil {
			usage()
		}
		countDays(nb, *z)
	} else {
		showTime(args)
	}
}

func stringHasSuffix(needle string, stack []string) (string, bool) {
	for _, v := range stack {
		if strings.HasSuffix(needle, v) {
			return v, true
		}
	}
	return "", false
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
		guessTimeFormat(args[0])
		return
	}
	date := guessTimestamp(int64(timestamp))
	printTime(date)
}

func guessTimeFormat(arg string) {
	var format string
	if strings.Contains(arg, " ") {
		format = "2006-01-02 15:04:05"
	} else {
		format = "2006-01-02"
	}

	t, err := time.Parse(format, arg)
	if err != nil {
		fmt.Printf("Error: unknown format \"%s\" \n", arg)
		usage()
		return
	}

	printTime(t)
	return
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
	-z	Make the timestamp at 0 hour, 0 minute, 0 second of that day

TIMESTAMP:
	Timestamp can be in seconds, millioseconds, microseconds or nanoseconds

EXPR:
	Expressions can be shortcut words:
		tomorrow		timestamp 1 day in the future
		yesterday		timestamp 1 day in the past
		2daysago		timestamp 2 days in the past
		5daysfwd		timestamp 5 days in the future
		2014-12-24		timestamp representation of the date

`

func usage() {
	fmt.Println(usageinfo)
	os.Exit(2)
}
