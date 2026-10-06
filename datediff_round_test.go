package datediff

import (
	"math"
	"testing"
)

func TestRound(t *testing.T) {
	cases := []struct {
		v    float64
		n    int
		want float64
	}{
		{0, 0, 0},
		{1.4, 0, 1},
		{1.5, 0, 2},
		{1.6, 0, 2},
		{-1.4, 0, -1},
		{-1.5, 0, -2}, // HALF_UP: 负数 0.5 向绝对值大方向
		{-1.6, 0, -2},
		{1.234, 2, 1.23},
		{1.235, 2, 1.24},
		{1.245, 2, 1.25}, // 注意：1.245 可精确表示时的行为
		{-1.235, 2, -1.24},
		{3.1415926, 4, 3.1416},
		{0.0, 2, 0},
		{100, 2, 100},
		{1234.567, -2, 1200},  // n 为负：保留到百位
		{1250, -2, 1300},      // 1250 百位四舍五入
		{-1250, -2, -1300},
		{1.005, 2, 1.0},       // 经典浮点尾差：1.005 实际略小于 1.005，可能得 1.00
		{0.0, 10, 0},
	}
	for _, c := range cases {
		got := Round(c.v, c.n)
		if math.Abs(got-c.want) > 1e-9 {
			t.Errorf("Round(%v, %d) = %v want %v", c.v, c.n, got, c.want)
		}
	}
}

func TestRoundSpecial(t *testing.T) {
	if !math.IsNaN(Round(math.NaN(), 2)) {
		t.Error("NaN should return NaN")
	}
	if !math.IsInf(Round(math.Inf(1), 2), 1) {
		t.Error("+Inf should return +Inf")
	}
	if !math.IsInf(Round(math.Inf(-1), 2), -1) {
		t.Error("-Inf should return -Inf")
	}
}

func TestRoundMethod(t *testing.T) {
	d := &Datetime{}
	if got := d.Round(1.235, 2); math.Abs(got-1.24) > 1e-9 {
		t.Errorf("Datetime.Round(1.235,2) = %v want 1.24", got)
	}
}
