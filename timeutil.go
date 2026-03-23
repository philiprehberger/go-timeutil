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

// StartOfQuarter returns the first day of the quarter at 00:00:00.000000000,
// preserving the original time.Location. Quarters: Q1=Jan-Mar, Q2=Apr-Jun,
// Q3=Jul-Sep, Q4=Oct-Dec.
func StartOfQuarter(t time.Time) time.Time {
	y, m, _ := t.Date()
	quarterMonth := m - (m-1)%3
	return time.Date(y, quarterMonth, 1, 0, 0, 0, 0, t.Location())
}

// EndOfQuarter returns the last day of the quarter at 23:59:59.999999999,
// preserving the original time.Location.
func EndOfQuarter(t time.Time) time.Time {
	y, m, _ := t.Date()
	quarterMonth := m - (m-1)%3
	lastMonthOfQuarter := quarterMonth + 2
	firstOfNext := time.Date(y, lastMonthOfQuarter+1, 1, 0, 0, 0, 0, t.Location())
	lastDay := firstOfNext.AddDate(0, 0, -1)
	return EndOfDay(lastDay)
}

// StartOfYear returns January 1 of t's year at 00:00:00.000000000,
// preserving the original time.Location.
func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear returns December 31 of t's year at 23:59:59.999999999,
// preserving the original time.Location.
func EndOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 12, 31, 23, 59, 59, 999999999, t.Location())
}

// PreviousBusinessDay returns the start of the most recent weekday before t.
// If t is Monday, it returns the previous Friday. The original time.Location
// is preserved.
func PreviousBusinessDay(t time.Time) time.Time {
	prev := t.AddDate(0, 0, -1)
	for IsWeekend(prev) {
		prev = prev.AddDate(0, 0, -1)
	}
	return StartOfDay(prev)
}

// AddBusinessDays adds n business days (weekdays) to t. If n is negative,
// business days are subtracted. Weekends are skipped. The original
// time.Location is preserved.
func AddBusinessDays(t time.Time, n int) time.Time {
	direction := 1
	remaining := n
	if n < 0 {
		direction = -1
		remaining = -n
	}
	result := t
	for remaining > 0 {
		result = result.AddDate(0, 0, direction)
		if !IsWeekend(result) {
			remaining--
		}
	}
	return StartOfDay(result)
}

// BusinessDaysBetween returns the number of weekdays between a and b.
// The count is exclusive of the start date and inclusive of the end date.
// The result is always non-negative regardless of argument order.
func BusinessDaysBetween(a, b time.Time) int {
	a = StartOfDay(a)
	b = StartOfDay(b)
	if a.After(b) {
		a, b = b, a
	}
	count := 0
	cursor := a.AddDate(0, 0, 1)
	for !cursor.After(b) {
		if !IsWeekend(cursor) {
			count++
		}
		cursor = cursor.AddDate(0, 0, 1)
	}
	return count
}
