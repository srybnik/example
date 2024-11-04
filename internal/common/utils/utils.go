package utils

import (
	"time"
)

// Bow получает начало недели t
func Bow(t time.Time) time.Time {
	// Вычисляем разницу в днях между днем недели и началом недели
	daysSinceMonday := int(t.Weekday() - time.Monday)

	// Если текущий день недели - понедельник, то возвращаем ту же дату
	if daysSinceMonday == 0 {
		return Bod(t)
	}

	// Для случая, когда начало недели - воскресенье
	if daysSinceMonday < 0 {
		daysSinceMonday += 7
	}

	// Вычитаем разницу дней из текущей даты, чтобы получить начало недели
	return Bod(t.AddDate(0, 0, -daysSinceMonday))
}

// Bod получает начало дня t
func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// Eod получает конец дня t
func Eod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}

func SplitSlice[T any](s []T, count int) [][]T {
	var (
		res [][]T
		i   = 0
		j   = min(count, len(s))
	)
	for i < j {
		res = append(res, s[i:j])
		i = j
		j = min(j+count, len(s))
	}
	return res
}
