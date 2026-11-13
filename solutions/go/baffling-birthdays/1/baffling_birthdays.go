package bafflingbirthdays

import (
	"math"
	"math/rand"
	"time"
)

func SharedBirthday(dates []time.Time) bool {
	for i := 0; i < len(dates)-1; i++ {
		for j := i + 1; j < len(dates); j++ {
			if dates[i].Month() == dates[j].Month() && dates[i].Day() == dates[j].Day() {
				return true
			}
		}
	}
	return false
}

func RandomBirthdates(size int) []time.Time {
	dates := []time.Time{}
	for i := 0; i < size; i++ {
		date := time.Date(1999, time.January, 1, 0, 0, 0, 0, time.UTC)
		date = date.AddDate(0, 0, rand.Intn(365))
		dates = append(dates, date)
	}
	return dates
}

func EstimatedProbability(size int) float64 {
	return (1 - math.Pow(math.E, -(math.Pow(float64(size), 2)/730))) * 100
}
