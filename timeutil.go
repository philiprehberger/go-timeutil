// Package timeutil provides time manipulation helpers for Go.
package timeutil

import "time"

// StartOfDay returns midnight (00:00:00.000000000) of the given day,
// preserving the original time.Location.
func StartOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// EndOfDay returns the last nanosecond (23:59:59.999999999) of the given day,
// preserving the original time.Location.
func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, 999999999, t.Location())
}

// StartOfWeek returns Monday 00:00:00.000000000 of the week containing t.
// Weeks start on Monday (ISO 8601). The original time.Location is preserved.
func StartOfWeek(t time.Time) time.Time {
	weekday := t.Weekday()
	if weekday == time.Sunday {
		weekday = 7
	}
	offset := int(weekday) - int(time.Monday)
	monday := t.AddDate(0, 0, -offset)
	return StartOfDay(monday)
}

// EndOfWeek returns Sunday 23:59:59.999999999 of the week containing t.
// Weeks start on Monday (ISO 8601). The original time.Location is preserved.
func EndOfWeek(t time.Time) time.Time {
	weekday := t.Weekday()
	if weekday == time.Sunday {
		return EndOfDay(t)
	}
	daysUntilSunday := 7 - int(weekday)
	sunday := t.AddDate(0, 0, daysUntilSunday)
	return EndOfDay(sunday)
}

// StartOfMonth returns the first day of the month at 00:00:00.000000000,
// preserving the original time.Location.
func StartOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth returns the last day of the month at 23:59:59.999999999,
// preserving the original time.Location.
func EndOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	// First day of next month, then subtract one day.
	firstOfNext := time.Date(y, m+1, 1, 0, 0, 0, 0, t.Location())
	lastDay := firstOfNext.AddDate(0, 0, -1)
	return EndOfDay(lastDay)
}

// DaysBetween returns the absolute number of calendar days between a and b.
// It compares dates only, ignoring the time-of-day component.
func DaysBetween(a, b time.Time) int {
	a = StartOfDay(a)
	b = StartOfDay(b)
	diff := b.Sub(a)
	days := int(diff.Hours() / 24)
	if days < 0 {
		days = -days
	}
	return days
}

// IsWeekend reports whether t falls on Saturday or Sunday.
func IsWeekend(t time.Time) bool {
	day := t.Weekday()
	return day == time.Saturday || day == time.Sunday
}

// IsBusinessDay reports whether t falls on a weekday (Monday through Friday).
func IsBusinessDay(t time.Time) bool {
	return !IsWeekend(t)
}

// NextBusinessDay returns the start of the next weekday (Monday through Friday)
// after t. If t is a weekday, it returns the following weekday. If t falls on
// Friday, it returns the following Monday. The original time.Location is preserved.
func NextBusinessDay(t time.Time) time.Time {
	next := t.AddDate(0, 0, 1)
	for IsWeekend(next) {
		next = next.AddDate(0, 0, 1)
	}
	return StartOfDay(next)
}
