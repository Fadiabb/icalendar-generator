package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/anders/ics"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func getEnterdValues() (x, y string) {
	var date, recordsNumber string
	scanner := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter starting date in form yyyy-mm-dd")
		date, _ = scanner.ReadString('\n')
		date = strings.TrimSuffix(date, "\n")
		match, _ := regexp.MatchString(`^\d{4}-([0-1]{1})?[0-9]{1}-\d{1,2}`, date)
		if match {
			break
		}

	}
	for {
		fmt.Println("Enter Number of records")
		fmt.Scanln(&recordsNumber)
		match, _ := regexp.MatchString(`\d`, recordsNumber)
		if match {
			break
		}
	}
	return date, recordsNumber
}
func main() {
	date, recordsNumber := getEnterdValues()
	number, _ := strconv.Atoi(recordsNumber)
	startingDate, err := time.Parse("2006-1-2", date)
	check(err)
	cal := ics.NewCalendar()
	for i := 0; i <= number; i++ {
		cal.Add(ics.Event{
			"DTSTART": startingDate,
			/* "DTEND": endingDate, */
			"SUMMARY": "Event",
		})
	}
	fp, err := os.Create("calender.ics")
	if err != nil {
		log.Fatalln("error by creating .ics file")
		return
	}
	defer fp.Close()
	cal.Encode(fp)
	fmt.Println("calender.ics file generated successfully")
}
