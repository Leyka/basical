package utils

import (
	"fmt"
	"time"
)

var daysFR = map[string]string{
	"Monday":    "Lundi",
	"Tuesday":   "Mardi",
	"Wednesday": "Mercredi",
	"Thursday":  "Jeudi",
	"Friday":    "Vendredi",
	"Saturday":  "Samedi",
	"Sunday":    "Dimanche",
}

var monthsFR = map[string]string{
	"January":   "Janvier",
	"February":  "Février",
	"March":     "Mars",
	"April":     "Avril",
	"May":       "Mai",
	"June":      "Juin",
	"July":      "Juillet",
	"August":    "Août",
	"September": "Septembre",
	"October":   "Octobre",
	"November":  "Novembre",
	"December":  "Décembre",
}

// Returns the formatted date in French
// Example: "Lundi 1 Janvier"
func GetFormattedDateInFrench(t time.Time) string {
	dayName := t.Weekday().String()
	dayFr := daysFR[dayName]

	monthName := t.Month().String()
	monthFr := monthsFR[monthName]

	return fmt.Sprintf(
		"%s %d %s",
		dayFr,
		t.Day(),
		monthFr,
	)
}
