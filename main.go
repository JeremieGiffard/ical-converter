package main

import (
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

	columnsSlice := []string{"dtstart", "dtend", "description", "location", "summary"}

}

func isCallendarValid(sliceStr []string) bool {
	var isValid = false
	fmt.Println("\n isCallendarValid()", " sliceStr[0] =", sliceStr[0], " sliceStr[last] =", sliceStr[len(sliceStr)-1])
	if sliceStr[0] == "BEGIN:VCALENDAR" && sliceStr[len(sliceStr)-1] == "END:VCALENDAR" {
		isValid = true
	}

	return isValid
}
