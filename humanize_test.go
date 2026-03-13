package timeutil

import (
	"testing"
	"time"
)

func TestHumanizeJustNow(t *testing.T) {
	now := time.Now()
	result := humanizeRelative(now, now)
	if result != "just now" {
		t.Errorf("got %q, want %q", result, "just now")
	}

	result2 := humanizeRelative(now.Add(-2*time.Second), now)
	if result2 != "just now" {
		t.Errorf("got %q, want %q", result2, "just now")
	}
}

func TestHumanizePast(t *testing.T) {
	now := time.Date(2026, 3, 15, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		t    time.Time
		want string
	}{
		{"seconds ago", now.Add(-30 * time.Second), "30 seconds ago"},
		{"1 minute ago", now.Add(-90 * time.Second), "2 minutes ago"},
		{"minutes ago", now.Add(-5 * time.Minute), "5 minutes ago"},
		{"1 hour ago", now.Add(-1 * time.Hour), "1 hour ago"},
		{"hours ago", now.Add(-3 * time.Hour), "3 hours ago"},
		{"yesterday", now.Add(-30 * time.Hour), "yesterday"},
		{"days ago", now.Add(-72 * time.Hour), "3 days ago"},
		{"weeks ago", now.Add(-14 * 24 * time.Hour), "2 weeks ago"},
		{"1 month ago", now.Add(-35 * 24 * time.Hour), "1 month ago"},
		{"months ago", now.Add(-90 * 24 * time.Hour), "3 months ago"},
		{"1 year ago", now.Add(-400 * 24 * time.Hour), "1 year ago"},
		{"years ago", now.Add(-800 * 24 * time.Hour), "2 years ago"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := humanizeRelative(tt.t, now)
			if result != tt.want {
				t.Errorf("got %q, want %q", result, tt.want)
			}
		})
	}
}

func TestHumanizeFuture(t *testing.T) {
	now := time.Date(2026, 3, 15, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		t    time.Time
		want string
	}{
		{"in seconds", now.Add(30 * time.Second), "in 30 seconds"},
		{"in minutes", now.Add(5 * time.Minute), "in 5 minutes"},
		{"in hours", now.Add(3 * time.Hour), "in 3 hours"},
		{"in days", now.Add(72 * time.Hour), "in 3 days"},
		{"in weeks", now.Add(14 * 24 * time.Hour), "in 2 weeks"},
		{"in months", now.Add(90 * 24 * time.Hour), "in 3 months"},
		{"in years", now.Add(800 * 24 * time.Hour), "in 2 years"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := humanizeRelative(tt.t, now)
			if result != tt.want {
				t.Errorf("got %q, want %q", result, tt.want)
			}
		})
	}
}

func TestHumanizeDuration(t *testing.T) {
	tests := []struct {
		name string
		d    time.Duration
		want string
	}{
		{"zero", 0, "0 seconds"},
		{"seconds", 45 * time.Second, "45 seconds"},
		{"1 second", 1 * time.Second, "1 second"},
		{"minutes and seconds", 3*time.Minute + 15*time.Second, "3 minutes 15 seconds"},
		{"hours and minutes", 2*time.Hour + 30*time.Minute, "2 hours 30 minutes"},
		{"days and hours", 26*time.Hour + 30*time.Minute, "1 day 2 hours 30 minutes"},
		{"exact hours", 5 * time.Hour, "5 hours"},
		{"1 day", 24 * time.Hour, "1 day"},
		{"complex", 49*time.Hour + 5*time.Minute + 10*time.Second, "2 days 1 hour 5 minutes 10 seconds"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HumanizeDuration(tt.d)
			if result != tt.want {
				t.Errorf("got %q, want %q", result, tt.want)
			}
		})
	}
}

func TestHumanizeDurationNegative(t *testing.T) {
	result := HumanizeDuration(-2*time.Hour - 30*time.Minute)
	want := "-2 hours 30 minutes"
	if result != want {
		t.Errorf("got %q, want %q", result, want)
	}

	result2 := HumanizeDuration(-45 * time.Second)
	want2 := "-45 seconds"
	if result2 != want2 {
		t.Errorf("got %q, want %q", result2, want2)
	}
}
