package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	sliceStr, _ := readFileAndSlice("callendar-example.ics")
	isCallendarValid(sliceStr)

	SliceColumn := []string{"dtstart", "dtend", "description", "location", "summary"}

	CsvMapFromIcs := makeCsvMapFromIcs(sliceStr)
	// fmt.Println(CsvMapFromIcs)
	ColumnsNames, ValuesToColumns := filterColumn(CsvMapFromIcs, SliceColumn)

	csvData := [][]string{
		ColumnsNames, ValuesToColumns}

	// for _, value := range ValuesToColumns {
	// 	csvData = append(csvData, []string{value})
	// }

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

func filterColumn(mapStr map[string]string, ColumnsWanted []string) ([]string, []string) {
	keysToReturn := []string{}
	valuesToReturn := []string{}
	for _, key := range ColumnsWanted {
		value, keyExist := mapStr[strings.ToUpper(key)]
		if keyExist {
			keysToReturn = append(keysToReturn, key)
			valuesToReturn = append(valuesToReturn, value)
		}
		fmt.Println("After filter key ", keysToReturn)
		fmt.Println("After filter value ", valuesToReturn)
	}
	return keysToReturn, valuesToReturn
}

func readFileAndSlice(filePath string) ([]string, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	str := string(b)
	sliceStr := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")
	//fmt.Println("\n\n\n\nslice converted", sliceStr)
	return sliceStr, err
}

func isCallendarValid(sliceStr []string) bool {
	var isValid = false
	fmt.Println("\n isCallendarValid()", " sliceStr[0] =", sliceStr[0], " sliceStr[last] =", sliceStr[len(sliceStr)-1])
	if sliceStr[0] == "BEGIN:VCALENDAR" && sliceStr[len(sliceStr)-1] == "END:VCALENDAR" {
		isValid = true
	}

	return isValid
}

func makeCsvMapFromIcs(sliceStr []string) map[string]string {
	var csvDataTable = make(map[string]string)
	for _, v := range sliceStr {
		separator := ":"
		if strings.Contains(v, "DTSTART;") || strings.Contains(v, "DTEND;") || strings.Contains(v, "ATTACH;") {
			separator = ";"
		}

		SliceKeyAndValue := strings.SplitN(v, separator, 2)
		csvDataTable[SliceKeyAndValue[0]] = SliceKeyAndValue[1]

	}
	fmt.Printf("%q\n", csvDataTable)
	return csvDataTable
}
