package datediff

import (
	"fmt"
	"testing"
	"time"
)

// 验证 Birthday 时分秒敏感性已修复、Days 改用 Local 后行为一致
func TestLocalTimeNormalize(t *testing.T) {
	d := &Datetime{}

	fmt.Println("--- Birthday: 生日 1/1 12:00，今天 1/1 09:00 应得 1 周年 ---")
	b := time.Date(2000, 1, 1, 12, 0, 0, 0, time.Local)
	e := time.Date(2001, 1, 1, 9, 0, 0, 0, time.Local)
	y, m, dd := d.Birthday(b, e)
	if y != 1 || m != 0 || dd != 0 {
		t.Errorf("got (%d,%d,%d) want (1,0,0)", y, m, dd)
	}
	fmt.Printf("  -> y=%d m=%d d=%d ✓\n", y, m, dd)

	fmt.Println("--- Birthday: 同时刻对比 ---")
	e2 := time.Date(2001, 1, 1, 12, 0, 0, 0, time.Local)
	y, m, dd = d.Birthday(b, e2)
	if y != 1 || m != 0 || dd != 0 {
		t.Errorf("got (%d,%d,%d) want (1,0,0)", y, m, dd)
	}
	fmt.Printf("  -> y=%d m=%d d=%d ✓\n", y, m, dd)

	fmt.Println("--- Days: 时分秒敏感已修复 ---")
	b3 := time.Date(2026, 8, 24, 23, 0, 0, 0, time.Local)
	e3 := time.Date(2026, 10, 5, 1, 0, 0, 0, time.Local)
	got := d.Days(b3, e3)
	if got != 42 {
		t.Errorf("got %d want 42", got)
	}
	fmt.Printf("  Days(8/24 23:00, 10/5 01:00) = %d ✓\n", got)
}
