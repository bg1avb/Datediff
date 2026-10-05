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
	//归一到本地时区午夜，避免 b、e 携带的时分秒影响日历日比较
	//(例如 b=12:00 而 e=time.Now() 上午取值会少算 1 天)。
	b = time.Date(b.Year(), b.Month(), b.Day(), 0, 0, 0, 0, time.Local)
	e = time.Date(e.Year(), e.Month(), e.Day(), 0, 0, 0, 0, time.Local)
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
	//仅 2/29 生日跨年到非闰年时，Go 的 AddDate 把周年日溢出到 3/1。
	//此处按 Feb 28 截断约定补偿 1 天。其他月末生日（1/31、3/31 等）
	//的溢出属于 Go 的前向归一化，不应补偿，否则会多算 1 天。
	if b.Month() == time.February && b.Day() == 29 && closeDay.Day() != b.Day() {
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

// Days 返回两个日期之间的纯天数，结果非负（b 在后则自动交换）。
// 用于新生儿期等需要"满 N 天"的场景。
// 与 Birthday 不同，不会用月份填满，因此跨月/跨闰年时
// 仍只返回天数；不能复用 Birthday 的 day 字段。
// 内部归一到本地时区午夜，与 Birthday 保持一致。
func (*Datetime) Days(b, e time.Time) int {
	if b.After(e) {
		b, e = e, b
	}
	b = time.Date(b.Year(), b.Month(), b.Day(), 0, 0, 0, 0, time.Local)
	e = time.Date(e.Year(), e.Month(), e.Day(), 0, 0, 0, 0, time.Local)
	n := 0
	for b.AddDate(0, 0, n).Before(e) {
		n++
	}
	return n
}
