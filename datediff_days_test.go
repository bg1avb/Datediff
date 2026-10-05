package datediff

import (
	"fmt"
	"testing"
	"time"
)

func TestDays(t *testing.T) {
	d := &Datetime{}
	cases := []struct {
		b, e string
		want int
	}{
		{"2026-08-24", "2026-10-05", 42},   // 产后 42 天复查
		{"2020-02-28", "2020-03-01", 2},     // 跨闰年 2/29
		{"2019-02-28", "2019-03-01", 1},     // 平年对照
		{"2020-01-01", "2020-03-01", 60},    // 闰年 1/1→3/1
		{"2019-01-01", "2019-03-01", 59},    // 平年 1/1→3/1
		{"2020-02-29", "2020-02-29", 0},     // 同日
		{"2020-12-31", "2021-01-01", 1},     // 跨年
		{"2026-10-05", "2026-08-24", 42},    // 反向输入应得非负
	}
	for _, c := range cases {
		b, _ := time.Parse("2006-01-02", c.b)
		e, _ := time.Parse("2006-01-02", c.e)
		got := d.Days(b, e)
		if got != c.want {
			t.Errorf("Days(%s,%s)=%d want %d", c.b, c.e, got, c.want)
		}
		fmt.Printf("Days(%s,%s)=%d want=%d\n", c.b, c.e, got, c.want)
	}
}
