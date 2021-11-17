package datediff

import (
	"fmt"
	"testing"
	"time"
)

func TestAddBirthday(t *testing.T) {
	datetime := &Datetime{}
	birthday := time.Date(2000, 2, 29, 0, 0, 0, 0, time.Local)
	calcdate := time.Date(2000, 3, 28, 0, 0, 0, 0, time.Local)
	year, month, day := datetime.Birthday(birthday, calcdate)
	fmt.Println("Year:", year, " Month:", month, " day:", day)

}
