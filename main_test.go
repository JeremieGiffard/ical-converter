package main

import "testing"

func TestReadFileAndSlice(t *testing.T) {

	got, _ := readFileAndSlice("callendar-example.ics")

	if got == nil {
		t.Errorf("got nil value,  expected []string")
	}

}
func TestReadFileAndSliceKO(t *testing.T) {

	_, err := readFileAndSlice("callendar-example.ic")

	if err == nil {
		t.Errorf("wrong path,  expected err,  got %v", err)
	}

}
func TestIcalIsValidKO(t *testing.T) {
	var invalidInput = []string{"BEGIN:VCALENDAR", "END:BADVALUE"}
	got := isCallendarValid(invalidInput)

	if got != false {
		t.Errorf("got %v bool,  want false", got)
	}

}
func TestIcalIsValidOK(t *testing.T) {
	var validInput = []string{"BEGIN:VCALENDAR", "END:VCALENDAR"}
	got := isCallendarValid(validInput)

	if got != true {
		t.Errorf("got %v bool,  want true", got)
	}

}
