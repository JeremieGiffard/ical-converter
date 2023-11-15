package main

import (
	"fmt"
	"os"

	ics "github.com/arran4/golang-ical"
)

func main() {
	b, _ := os.Open("callendar-example.ics")
	cal, err := ics.ParseCalendar(b)
	fmt.Println(cal, "/n")
	fmt.Println(err)
	// sliceStr, _ := readFileAndSlice("callendar-example.ics")
	// isCallendarValid(sliceStr)

	// SliceColumn := []string{"Start Date", "End Date", "description", "location", "summary"}

	// CsvMapFromIcs := makeCsvMapFromIcs(sliceStr)
	// // fmt.Println(CsvMapFromIcs)
	// ColumnsNames, ValuesToColumns := filterColumn(CsvMapFromIcs, SliceColumn)

	// csvData := [][]string{
	// 	ColumnsNames, ValuesToColumns}

	// // for _, value := range ValuesToColumns {
	// // 	csvData = append(csvData, []string{value})
	// // }

	// file, e := os.Create("./callendar-converted.csv")
	// if e != nil {
	// 	fmt.Println(e)
	// }
	// writer := csv.NewWriter(file)
	// e = writer.WriteAll(csvData)
	// if e != nil {
	// 	fmt.Println(e)
	// }

}

// func readFileAndSlice(filePath string) ([]string, error) {
// 	b, err := os.ReadFile(filePath)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	str := string(b)
// 	if isCallendarValid(str) {

// 		sliceStr := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")
// 	} else {
// 		panic(fmt.Sprintf("File not a valid iCalendar"))
// 	}

// 	//fmt.Println("\n\n\n\nslice converted", sliceStr)
// 	return sliceStr, err
// }

// func isCallendarValid(str string) bool {
// 	var isValid = false

// 	if strings.Contains(str, "BEGIN:VCALENDAR") && strings.Contains(str, "END:VCALENDAR") {
// 		isValid = true
// 	}
// 	return isValid
// }
// func makeCsvMapFromIcs(sliceStr []string) map[string]string {
// 	fmt.Printf("sliceStr All Data %q\n", sliceStr)
// 	var csvDataTable = make(map[string]string)
// 	for _, v := range sliceStr {
// 		separator := ":"
// 		if strings.Contains(v, "DTSTART;") || strings.Contains(v, "DTEND;") || strings.Contains(v, "ATTACH;") {
// 			separator = ";"
// 		}
// 		SliceKeyAndValue := strings.SplitN(v, separator, 2)

// 		if strings.Contains(SliceKeyAndValue[1], "TZID") {
// 			csvDataTable = handleDateTimeFormat(SliceKeyAndValue)

// 		} else {
// 			csvDataTable[SliceKeyAndValue[0]] = SliceKeyAndValue[1]
// 		}

// 	}
// 	fmt.Printf("%q\n", csvDataTable)
// 	return csvDataTable
// }
// func handleDateTimeFormat(SliceDateTime []string) map[string]string {
// 	var csvDataTable = make(map[string]string)
// 	//example Date data : TZID=Europe/Paris:20230926T210000
// 	fmt.Printf("SliceDateTime %q\n", SliceDateTime)
// 	tempSlice := strings.SplitN(SliceDateTime[1], ":", 2)

// 	var DateAndTime []string
// 	if strings.Contains(tempSlice[1], "T") {
// 		//example Date data : 20230926T210000
// 		DateAndTime = strings.SplitN(tempSlice[1], "T", 2)
// 		fmt.Printf("DateAndTime %q\n", DateAndTime)
// 	}

// 	if SliceDateTime[0] == "DTSTART" {
// 		csvDataTable["Start Date"] = DateAndTime[0]
// 		csvDataTable["Start Time"] = DateAndTime[1]
// 	} else if SliceDateTime[0] == "DTEND" {
// 		csvDataTable["End Date"] = DateAndTime[0]
// 		csvDataTable["End Time"] = DateAndTime[1]
// 	}
// 	fmt.Printf("csvDataTable %q\n", csvDataTable)
// 	return csvDataTable
// }

// func filterColumn(mapStr map[string]string, ColumnsWanted []string) ([]string, []string) {
// 	keysToReturn := []string{}
// 	valuesToReturn := []string{}
// 	fmt.Printf("mapStr %q\n", mapStr)
// 	for _, key := range ColumnsWanted {
// 		value, keyExist := mapStr[strings.ToUpper(key)]
// 		if keyExist {
// 			keysToReturn = append(keysToReturn, key)
// 			valuesToReturn = append(valuesToReturn, value)
// 		}
// 		// fmt.Println("After filter key ", keysToReturn)
// 		// fmt.Println("After filter value ", valuesToReturn)
// 	}
// 	return keysToReturn, valuesToReturn
// }
