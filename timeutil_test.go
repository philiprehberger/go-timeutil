package timeutil

import (
	"testing"
	"time"
)

func TestStartOfDay(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	input := time.Date(2026, 3, 15, 14, 30, 45, 123456789, loc)
	result := StartOfDay(input)

	if result.Year() != 2026 || result.Month() != 3 || result.Day() != 15 {
		t.Errorf("date changed: got %v", result)
	}
	if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 || result.Nanosecond() != 0 {
		t.Errorf("expected midnight, got %v", result)
	}
	if result.Location() != loc {
		t.Errorf("location changed: got %v, want %v", result.Location(), loc)
	}
}

func TestEndOfDay(t *testing.T) {
	input := time.Date(2026, 7, 4, 10, 0, 0, 0, time.UTC)
	result := EndOfDay(input)

	if result.Year() != 2026 || result.Month() != 7 || result.Day() != 4 {
		t.Errorf("date changed: got %v", result)
	}
	if result.Hour() != 23 || result.Minute() != 59 || result.Second() != 59 || result.Nanosecond() != 999999999 {
		t.Errorf("expected 23:59:59.999999999, got %02d:%02d:%02d.%09d",
			result.Hour(), result.Minute(), result.Second(), result.Nanosecond())
	}
}

func TestStartOfWeek(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		wantDate time.Time
	}{
		{
			name:     "Monday stays Monday",
			input:    time.Date(2026, 3, 9, 15, 0, 0, 0, time.UTC), // Monday
			wantDate: time.Date(2026, 3, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Wednesday goes to Monday",
			input:    time.Date(2026, 3, 11, 12, 0, 0, 0, time.UTC), // Wednesday
			wantDate: time.Date(2026, 3, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Friday goes to Monday",
			input:    time.Date(2026, 3, 13, 18, 0, 0, 0, time.UTC), // Friday
			wantDate: time.Date(2026, 3, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Saturday goes to Monday",
			input:    time.Date(2026, 3, 14, 10, 0, 0, 0, time.UTC), // Saturday
			wantDate: time.Date(2026, 3, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Sunday goes to Monday",
			input:    time.Date(2026, 3, 15, 20, 0, 0, 0, time.UTC), // Sunday
			wantDate: time.Date(2026, 3, 9, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfWeek(tt.input)
			if !result.Equal(tt.wantDate) {
				t.Errorf("got %v, want %v", result, tt.wantDate)
			}
			if result.Weekday() != time.Monday {
				t.Errorf("expected Monday, got %v", result.Weekday())
			}
		})
	}
}

func TestEndOfWeek(t *testing.T) {
	// Wednesday March 11, 2026
	input := time.Date(2026, 3, 11, 12, 0, 0, 0, time.UTC)
	result := EndOfWeek(input)

	// Sunday March 15, 2026
	if result.Weekday() != time.Sunday {
		t.Errorf("expected Sunday, got %v", result.Weekday())
	}
	if result.Day() != 15 {
		t.Errorf("expected day 15, got %d", result.Day())
	}
	if result.Hour() != 23 || result.Minute() != 59 || result.Second() != 59 || result.Nanosecond() != 999999999 {
		t.Error("expected end of day")
	}

	// Sunday should return end of the same day
	sunday := time.Date(2026, 3, 15, 10, 0, 0, 0, time.UTC)
	result2 := EndOfWeek(sunday)
	if result2.Day() != 15 {
		t.Errorf("Sunday: expected day 15, got %d", result2.Day())
	}
}

func TestStartOfMonth(t *testing.T) {
	input := time.Date(2026, 8, 25, 14, 30, 0, 0, time.UTC)
	result := StartOfMonth(input)

	if result.Day() != 1 {
		t.Errorf("expected day 1, got %d", result.Day())
	}
	if result.Month() != 8 || result.Year() != 2026 {
		t.Errorf("month/year changed: got %v", result)
	}
	if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
		t.Error("expected midnight")
	}
}

func TestEndOfMonth(t *testing.T) {
	tests := []struct {
		name    string
		input   time.Time
		wantDay int
	}{
		{
			name:    "January has 31 days",
			input:   time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			wantDay: 31,
		},
		{
			name:    "February non-leap has 28 days",
			input:   time.Date(2026, 2, 10, 0, 0, 0, 0, time.UTC),
			wantDay: 28,
		},
		{
			name:    "February leap year has 29 days",
			input:   time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC),
			wantDay: 29,
		},
		{
			name:    "April has 30 days",
			input:   time.Date(2026, 4, 20, 0, 0, 0, 0, time.UTC),
			wantDay: 30,
		},
		{
			name:    "December has 31 days",
			input:   time.Date(2026, 12, 1, 0, 0, 0, 0, time.UTC),
			wantDay: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndOfMonth(tt.input)
			if result.Day() != tt.wantDay {
				t.Errorf("got day %d, want %d", result.Day(), tt.wantDay)
			}
			if result.Hour() != 23 || result.Minute() != 59 || result.Second() != 59 || result.Nanosecond() != 999999999 {
				t.Error("expected end of day")
			}
		})
	}
}

func TestDaysBetween(t *testing.T) {
	tests := []struct {
		name string
		a, b time.Time
		want int
	}{
		{
			name: "same day",
			a:    time.Date(2026, 3, 15, 8, 0, 0, 0, time.UTC),
			b:    time.Date(2026, 3, 15, 20, 0, 0, 0, time.UTC),
			want: 0,
		},
		{
			name: "adjacent days",
			a:    time.Date(2026, 3, 15, 0, 0, 0, 0, time.UTC),
			b:    time.Date(2026, 3, 16, 0, 0, 0, 0, time.UTC),
			want: 1,
		},
		{
			name: "one week",
			a:    time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC),
			b:    time.Date(2026, 3, 8, 0, 0, 0, 0, time.UTC),
			want: 7,
		},
		{
			name: "reversed order (absolute)",
			a:    time.Date(2026, 3, 20, 0, 0, 0, 0, time.UTC),
			b:    time.Date(2026, 3, 15, 0, 0, 0, 0, time.UTC),
			want: 5,
		},
		{
			name: "across months",
			a:    time.Date(2026, 1, 31, 0, 0, 0, 0, time.UTC),
			b:    time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC),
			want: 29,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DaysBetween(tt.a, tt.b)
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

func TestIsWeekend(t *testing.T) {
	saturday := time.Date(2026, 3, 14, 12, 0, 0, 0, time.UTC)
	sunday := time.Date(2026, 3, 15, 12, 0, 0, 0, time.UTC)
	monday := time.Date(2026, 3, 16, 12, 0, 0, 0, time.UTC)
	friday := time.Date(2026, 3, 13, 12, 0, 0, 0, time.UTC)

	if !IsWeekend(saturday) {
		t.Error("Saturday should be weekend")
	}
	if !IsWeekend(sunday) {
		t.Error("Sunday should be weekend")
	}
	if IsWeekend(monday) {
		t.Error("Monday should not be weekend")
	}
	if IsWeekend(friday) {
		t.Error("Friday should not be weekend")
	}
}

func TestIsBusinessDay(t *testing.T) {
	monday := time.Date(2026, 3, 9, 12, 0, 0, 0, time.UTC)
	wednesday := time.Date(2026, 3, 11, 12, 0, 0, 0, time.UTC)
	friday := time.Date(2026, 3, 13, 12, 0, 0, 0, time.UTC)
	saturday := time.Date(2026, 3, 14, 12, 0, 0, 0, time.UTC)
	sunday := time.Date(2026, 3, 15, 12, 0, 0, 0, time.UTC)

	if !IsBusinessDay(monday) {
		t.Error("Monday should be business day")
	}
	if !IsBusinessDay(wednesday) {
		t.Error("Wednesday should be business day")
	}
	if !IsBusinessDay(friday) {
		t.Error("Friday should be business day")
	}
	if IsBusinessDay(saturday) {
		t.Error("Saturday should not be business day")
	}
	if IsBusinessDay(sunday) {
		t.Error("Sunday should not be business day")
	}
}

func TestNextBusinessDay(t *testing.T) {
	tests := []struct {
		name    string
		input   time.Time
		wantDay int
		wantWD  time.Weekday
	}{
		{
			name:    "Wednesday to Thursday",
			input:   time.Date(2026, 3, 11, 14, 0, 0, 0, time.UTC),
			wantDay: 12,
			wantWD:  time.Thursday,
		},
		{
			name:    "Friday to Monday",
			input:   time.Date(2026, 3, 13, 17, 0, 0, 0, time.UTC),
			wantDay: 16,
			wantWD:  time.Monday,
		},
		{
			name:    "Saturday to Monday",
			input:   time.Date(2026, 3, 14, 10, 0, 0, 0, time.UTC),
			wantDay: 16,
			wantWD:  time.Monday,
		},
		{
			name:    "Sunday to Monday",
			input:   time.Date(2026, 3, 15, 10, 0, 0, 0, time.UTC),
			wantDay: 16,
			wantWD:  time.Monday,
		},
		{
			name:    "Thursday to Friday",
			input:   time.Date(2026, 3, 12, 9, 0, 0, 0, time.UTC),
			wantDay: 13,
			wantWD:  time.Friday,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NextBusinessDay(tt.input)
			if result.Day() != tt.wantDay {
				t.Errorf("got day %d, want %d", result.Day(), tt.wantDay)
			}
			if result.Weekday() != tt.wantWD {
				t.Errorf("got weekday %v, want %v", result.Weekday(), tt.wantWD)
			}
			if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
				t.Error("expected start of day")
			}
		})
	}
}

func TestStartOfQuarter(t *testing.T) {
	tests := []struct {
		name      string
		input     time.Time
		wantMonth time.Month
		wantDay   int
	}{
		{
			name:      "Q1 - January",
			input:     time.Date(2026, 1, 15, 14, 30, 0, 0, time.UTC),
			wantMonth: time.January,
			wantDay:   1,
		},
		{
			name:      "Q1 - March",
			input:     time.Date(2026, 3, 31, 23, 59, 0, 0, time.UTC),
			wantMonth: time.January,
			wantDay:   1,
		},
		{
			name:      "Q2 - April",
			input:     time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
			wantMonth: time.April,
			wantDay:   1,
		},
		{
			name:      "Q2 - June",
			input:     time.Date(2026, 6, 20, 12, 0, 0, 0, time.UTC),
			wantMonth: time.April,
			wantDay:   1,
		},
		{
			name:      "Q3 - July",
			input:     time.Date(2026, 7, 4, 9, 0, 0, 0, time.UTC),
			wantMonth: time.July,
			wantDay:   1,
		},
		{
			name:      "Q3 - September",
			input:     time.Date(2026, 9, 30, 18, 0, 0, 0, time.UTC),
			wantMonth: time.July,
			wantDay:   1,
		},
		{
			name:      "Q4 - October",
			input:     time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC),
			wantMonth: time.October,
			wantDay:   1,
		},
		{
			name:      "Q4 - December",
			input:     time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
			wantMonth: time.October,
			wantDay:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfQuarter(tt.input)
			if result.Month() != tt.wantMonth {
				t.Errorf("got month %v, want %v", result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("got day %d, want %d", result.Day(), tt.wantDay)
			}
			if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 || result.Nanosecond() != 0 {
				t.Error("expected midnight")
			}
		})
	}
}

func TestEndOfQuarter(t *testing.T) {
	tests := []struct {
		name      string
		input     time.Time
		wantMonth time.Month
		wantDay   int
	}{
		{
			name:      "Q1 ends March 31",
			input:     time.Date(2026, 2, 10, 0, 0, 0, 0, time.UTC),
			wantMonth: time.March,
			wantDay:   31,
		},
		{
			name:      "Q2 ends June 30",
			input:     time.Date(2026, 5, 15, 0, 0, 0, 0, time.UTC),
			wantMonth: time.June,
			wantDay:   30,
		},
		{
			name:      "Q3 ends September 30",
			input:     time.Date(2026, 8, 1, 0, 0, 0, 0, time.UTC),
			wantMonth: time.September,
			wantDay:   30,
		},
		{
			name:      "Q4 ends December 31",
			input:     time.Date(2026, 11, 20, 0, 0, 0, 0, time.UTC),
			wantMonth: time.December,
			wantDay:   31,
		},
		{
			name:      "year boundary - Q4 of previous year",
			input:     time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
			wantMonth: time.December,
			wantDay:   31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndOfQuarter(tt.input)
			if result.Month() != tt.wantMonth {
				t.Errorf("got month %v, want %v", result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("got day %d, want %d", result.Day(), tt.wantDay)
			}
			if result.Hour() != 23 || result.Minute() != 59 || result.Second() != 59 || result.Nanosecond() != 999999999 {
				t.Error("expected end of day")
			}
		})
	}
}

func TestStartOfYear(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		wantYear int
	}{
		{
			name:     "mid year",
			input:    time.Date(2026, 7, 15, 14, 30, 0, 0, time.UTC),
			wantYear: 2026,
		},
		{
			name:     "last day of year",
			input:    time.Date(2025, 12, 31, 23, 59, 59, 999999999, time.UTC),
			wantYear: 2025,
		},
		{
			name:     "first day of year",
			input:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			wantYear: 2026,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfYear(tt.input)
			if result.Year() != tt.wantYear {
				t.Errorf("got year %d, want %d", result.Year(), tt.wantYear)
			}
			if result.Month() != time.January || result.Day() != 1 {
				t.Errorf("expected Jan 1, got %v %d", result.Month(), result.Day())
			}
			if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 || result.Nanosecond() != 0 {
				t.Error("expected midnight")
			}
		})
	}
}

func TestEndOfYear(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		wantYear int
	}{
		{
			name:     "mid year",
			input:    time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC),
			wantYear: 2026,
		},
		{
			name:     "first day of year",
			input:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			wantYear: 2024,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndOfYear(tt.input)
			if result.Year() != tt.wantYear {
				t.Errorf("got year %d, want %d", result.Year(), tt.wantYear)
			}
			if result.Month() != time.December || result.Day() != 31 {
				t.Errorf("expected Dec 31, got %v %d", result.Month(), result.Day())
			}
			if result.Hour() != 23 || result.Minute() != 59 || result.Second() != 59 || result.Nanosecond() != 999999999 {
				t.Error("expected end of day")
			}
		})
	}
}

func TestPreviousBusinessDay(t *testing.T) {
	tests := []struct {
		name    string
		input   time.Time
		wantDay int
		wantWD  time.Weekday
	}{
		{
			name:    "Wednesday to Tuesday",
			input:   time.Date(2026, 3, 11, 14, 0, 0, 0, time.UTC), // Wednesday
			wantDay: 10,
			wantWD:  time.Tuesday,
		},
		{
			name:    "Monday to Friday",
			input:   time.Date(2026, 3, 9, 9, 0, 0, 0, time.UTC), // Monday
			wantDay: 6,
			wantWD:  time.Friday,
		},
		{
			name:    "Sunday to Friday",
			input:   time.Date(2026, 3, 15, 12, 0, 0, 0, time.UTC), // Sunday
			wantDay: 13,
			wantWD:  time.Friday,
		},
		{
			name:    "Saturday to Friday",
			input:   time.Date(2026, 3, 14, 12, 0, 0, 0, time.UTC), // Saturday
			wantDay: 13,
			wantWD:  time.Friday,
		},
		{
			name:    "Tuesday to Monday",
			input:   time.Date(2026, 3, 10, 8, 0, 0, 0, time.UTC), // Tuesday
			wantDay: 9,
			wantWD:  time.Monday,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PreviousBusinessDay(tt.input)
			if result.Day() != tt.wantDay {
				t.Errorf("got day %d, want %d", result.Day(), tt.wantDay)
			}
			if result.Weekday() != tt.wantWD {
				t.Errorf("got weekday %v, want %v", result.Weekday(), tt.wantWD)
			}
			if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
				t.Error("expected start of day")
			}
		})
	}
}

func TestAddBusinessDays(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		n        int
		wantDate time.Time
	}{
		{
			name:     "add 1 from Wednesday",
			input:    time.Date(2026, 3, 11, 14, 0, 0, 0, time.UTC), // Wed
			n:        1,
			wantDate: time.Date(2026, 3, 12, 0, 0, 0, 0, time.UTC), // Thu
		},
		{
			name:     "add 1 from Friday skips weekend",
			input:    time.Date(2026, 3, 13, 10, 0, 0, 0, time.UTC), // Fri
			n:        1,
			wantDate: time.Date(2026, 3, 16, 0, 0, 0, 0, time.UTC), // Mon
		},
		{
			name:     "add 5 is one business week",
			input:    time.Date(2026, 3, 9, 0, 0, 0, 0, time.UTC), // Mon
			n:        5,
			wantDate: time.Date(2026, 3, 16, 0, 0, 0, 0, time.UTC), // next Mon
		},
		{
			name:     "add 0 returns start of same day",
			input:    time.Date(2026, 3, 11, 14, 0, 0, 0, time.UTC),
			n:        0,
			wantDate: time.Date(2026, 3, 11, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "negative subtracts days",
			input:    time.Date(2026, 3, 11, 10, 0, 0, 0, time.UTC), // Wed
			n:        -1,
			wantDate: time.Date(2026, 3, 10, 0, 0, 0, 0, time.UTC), // Tue
		},
		{
			name:     "negative across weekend",
			input:    time.Date(2026, 3, 9, 10, 0, 0, 0, time.UTC), // Mon
			n:        -1,
			wantDate: time.Date(2026, 3, 6, 0, 0, 0, 0, time.UTC), // Fri
		},
		{
			name:     "add from Saturday",
			input:    time.Date(2026, 3, 14, 10, 0, 0, 0, time.UTC), // Sat
			n:        1,
			wantDate: time.Date(2026, 3, 16, 0, 0, 0, 0, time.UTC), // Mon
		},
		{
			name:     "add 10 spans two weeks",
			input:    time.Date(2026, 3, 9, 0, 0, 0, 0, time.UTC), // Mon
			n:        10,
			wantDate: time.Date(2026, 3, 23, 0, 0, 0, 0, time.UTC), // Mon two weeks later
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddBusinessDays(tt.input, tt.n)
			if !result.Equal(tt.wantDate) {
				t.Errorf("got %v, want %v", result, tt.wantDate)
			}
		})
	}
}

func TestBusinessDaysBetween(t *testing.T) {
	tests := []struct {
		name string
		a, b time.Time
		want int
	}{
		{
			name: "same day",
			a:    time.Date(2026, 3, 11, 8, 0, 0, 0, time.UTC),
			b:    time.Date(2026, 3, 11, 20, 0, 0, 0, time.UTC),
			want: 0,
		},
		{
			name: "adjacent weekdays",
			a:    time.Date(2026, 3, 11, 0, 0, 0, 0, time.UTC), // Wed
			b:    time.Date(2026, 3, 12, 0, 0, 0, 0, time.UTC), // Thu
			want: 1,
		},
		{
			name: "Friday to Monday",
			a:    time.Date(2026, 3, 13, 0, 0, 0, 0, time.UTC), // Fri
			b:    time.Date(2026, 3, 16, 0, 0, 0, 0, time.UTC), // Mon
			want: 1,
		},
		{
			name: "full business week Mon-Fri",
			a:    time.Date(2026, 3, 9, 0, 0, 0, 0, time.UTC),  // Mon
			b:    time.Date(2026, 3, 13, 0, 0, 0, 0, time.UTC), // Fri
			want: 4,
		},
		{
			name: "Mon to next Mon",
			a:    time.Date(2026, 3, 9, 0, 0, 0, 0, time.UTC),  // Mon
			b:    time.Date(2026, 3, 16, 0, 0, 0, 0, time.UTC), // next Mon
			want: 5,
		},
		{
			name: "reversed order gives same result",
			a:    time.Date(2026, 3, 16, 0, 0, 0, 0, time.UTC), // Mon
			b:    time.Date(2026, 3, 9, 0, 0, 0, 0, time.UTC),  // prev Mon
			want: 5,
		},
		{
			name: "two full weeks",
			a:    time.Date(2026, 3, 9, 0, 0, 0, 0, time.UTC),  // Mon
			b:    time.Date(2026, 3, 23, 0, 0, 0, 0, time.UTC), // Mon +2 weeks
			want: 10,
		},
		{
			name: "across year boundary",
			a:    time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC), // Wed
			b:    time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC),   // Fri
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BusinessDaysBetween(tt.a, tt.b)
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

func TestBoundaryPreservesLocation(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	input := time.Date(2026, 6, 15, 14, 30, 0, 0, loc)

	funcs := map[string]func(time.Time) time.Time{
		"StartOfDay":     StartOfDay,
		"EndOfDay":       EndOfDay,
		"StartOfWeek":    StartOfWeek,
		"EndOfWeek":      EndOfWeek,
		"StartOfMonth":   StartOfMonth,
		"EndOfMonth":     EndOfMonth,
		"StartOfQuarter": StartOfQuarter,
		"EndOfQuarter":   EndOfQuarter,
		"StartOfYear":    StartOfYear,
		"EndOfYear":      EndOfYear,
	}

	for name, fn := range funcs {
		t.Run(name, func(t *testing.T) {
			result := fn(input)
			if result.Location() != loc {
				t.Errorf("%s: location = %v, want %v", name, result.Location(), loc)
			}
		})
	}
}
