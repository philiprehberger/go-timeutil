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

func TestBoundaryPreservesLocation(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	input := time.Date(2026, 6, 15, 14, 30, 0, 0, loc)

	funcs := map[string]func(time.Time) time.Time{
		"StartOfDay":   StartOfDay,
		"EndOfDay":     EndOfDay,
		"StartOfWeek":  StartOfWeek,
		"EndOfWeek":    EndOfWeek,
		"StartOfMonth": StartOfMonth,
		"EndOfMonth":   EndOfMonth,
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
