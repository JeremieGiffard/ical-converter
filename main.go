package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	b, err := os.ReadFile("callendar-example.ics")
	if err != nil {
		fmt.Println(err)
	}
	str := string(b)
	sliceStr := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")
	fmt.Println("\n\n\n\nslice converted", sliceStr)
	isCallendarValid(sliceStr)

	csvData := [][]string{
		{"dtstart", "dtend", "description", "location", "summary"}}

	file, e := os.Create("./callendar-converted.csv")
	if e != nil {
		fmt.Println(e)
	}
	writer := csv.NewWriter(file)
	e = writer.WriteAll(csvData)
	if e != nil {
		fmt.Println(e)
	}

}

func isCallendarValid(sliceStr []string) bool {
	var isValid = false
	fmt.Println("\n isCallendarValid()", " sliceStr[0] =", sliceStr[0], " sliceStr[last] =", sliceStr[len(sliceStr)-1])
	if sliceStr[0] == "BEGIN:VCALENDAR" && sliceStr[len(sliceStr)-1] == "END:VCALENDAR" {
		isValid = true
	}

	return isValid
}
