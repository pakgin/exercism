package booking

import (
	"fmt"
	"strings"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	_time, err := time.Parse(
		"1/2/2006 15:04:05",
		date,
	)
	if err != nil {
		panic(err)
	}
	return _time
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	_time, err := time.Parse(
		"January 2, 2006 15:04:05",
		date,
	)
	if err != nil {
		panic(err)
	}
	return _time.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	time_arr := strings.Split(date, " ")

	_time, _ := time.Parse("15:04:05", time_arr[4])
	start, _ := time.Parse("15:04:05", "11:59:59")
	end, _ := time.Parse("15:04:05", "18:00:00")

	return _time.After(start) && _time.Before(end)
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	_time, _ := time.Parse(
		"1/2/2006 15:04:05",
		date,
	)
	return fmt.Sprintf(
		"You have an appointment on %s, %s %d, %d, at %d:%d.",
		_time.Weekday().String(), _time.Month().String(),
		_time.Day(), _time.Year(), _time.Hour(), _time.Minute(),
	)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	this_year := fmt.Sprintf("%d-09-15 00:00:00 UTC", time.Now().Year())
	_time, err := time.Parse(
		"2006-01-02 15:04:05 MST",
		this_year,
	)
	if err != nil {
		panic(err)
	}
	return _time
}
