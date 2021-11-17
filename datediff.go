package datediff

import (
	"time"
)

type Datetime struct{}

/*
datediff
para : birth date
       calculate date
return calculate date - birth date = year month day
*/

func (*Datetime) Birthday(b, e time.Time) (year, month, day int) {
	var closeDay time.Time
	var leapday int = 0 //leap year Feb 29 add 1 year will by March 1
	year, month, day = 0, 0, 0
	if b.After(e) {
		b, e = e, b
	}
	//try years
	for {
		if !b.AddDate(year, 0, 0).After(e) {
			year++
		} else {
			year--
			break
		}
	}
	//try months
	for {
		if !b.AddDate(0, month, 0).After(e) {
			month++
		} else {
			month--
			break
		}
	}
	//try days add all months before then calcdate
	closeDay = b.AddDate(0, month, 0)
	if closeDay.Day() != b.Day() {
		leapday = 1
	}
	month = month % 12

	//try days
	for {
		if !closeDay.AddDate(0, 0, day).After(e) {
			//fmt.Println("day:", closeDay.AddDate(0, 0, day))
			day++
		} else {
			day--
			break
		}
	}
	day += leapday
	return
}
