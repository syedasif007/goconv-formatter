package formatter

import (
	"fmt"
	"strconv"
)

// StrFloat ...
func StrFloat(f float64) string {
	return fmt.Sprintf("%.2f", f)
}

// StrInt ...
func StrInt(i int) string {
	return fmt.Sprintf("%d", i)
}

// CnvFloat ...
func CnvFloat(s string) float64 {
	f, e := strconv.ParseFloat(s, 64)

	if e != nil {
		f = 0.0
	}

	return f
}

// CnvInt64 ...
func CnvInt64(s string) int64 {
	i64, e := strconv.ParseInt(s, 10, 64)

	if e != nil {
		i64 = 0
	}

	return i64
}

// CnvInt ...
func CnvInt(s string) int {
	i, e := strconv.Atoi(s)

	if e != nil {
		i = 0
	}

	return i
}
