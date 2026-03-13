package timeutil

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// Humanize returns a human-readable relative time string for t compared to
// the current time. Past times produce strings like "5 seconds ago" or
// "yesterday". Future times produce strings like "in 3 minutes".
func Humanize(t time.Time) string {
	return humanizeRelative(t, time.Now())
}

// humanizeRelative is the internal implementation that accepts a reference time
// for testability.
func humanizeRelative(t, now time.Time) string {
	diff := now.Sub(t)
	if diff < 0 {
		absDiff := -diff
		days := int(math.Round(absDiff.Hours() / 24))
		if days == 1 {
			return "tomorrow"
		}
		return "in " + humanizeAbs(absDiff)
	}
	if diff < 5*time.Second {
		return "just now"
	}
	days := int(math.Round(diff.Hours() / 24))
	if days == 1 {
		return "yesterday"
	}
	result := humanizeAbs(diff)
	return result + " ago"
}

func humanizeAbs(d time.Duration) string {
	seconds := int(math.Round(d.Seconds()))
	if seconds < 60 {
		return pluralize(seconds, "second")
	}

	minutes := int(math.Round(d.Minutes()))
	if minutes < 60 {
		return pluralize(minutes, "minute")
	}

	hours := int(math.Round(d.Hours()))
	if hours < 24 {
		return pluralize(hours, "hour")
	}

	days := int(math.Round(d.Hours() / 24))
	if days < 7 {
		return pluralize(days, "day")
	}

	weeks := days / 7
	if weeks < 4 {
		return pluralize(weeks, "week")
	}

	months := days / 30
	if months < 12 {
		if months == 0 {
			months = 1
		}
		return pluralize(months, "month")
	}

	years := days / 365
	if years == 0 {
		years = 1
	}
	return pluralize(years, "year")
}

func pluralize(n int, unit string) string {
	if n == 1 {
		return fmt.Sprintf("1 %s", unit)
	}
	return fmt.Sprintf("%d %ss", n, unit)
}

// HumanizeDuration returns a human-readable string representation of a
// time.Duration, broken down into days, hours, minutes, and seconds.
// Zero-value components are omitted. Negative durations are prefixed with "-".
// A zero duration returns "0 seconds".
func HumanizeDuration(d time.Duration) string {
	if d == 0 {
		return "0 seconds"
	}

	prefix := ""
	if d < 0 {
		prefix = "-"
		d = -d
	}

	totalSeconds := int(d.Seconds())
	days := totalSeconds / 86400
	remaining := totalSeconds % 86400
	hours := remaining / 3600
	remaining = remaining % 3600
	minutes := remaining / 60
	seconds := remaining % 60

	var parts []string
	if days > 0 {
		parts = append(parts, pluralize(days, "day"))
	}
	if hours > 0 {
		parts = append(parts, pluralize(hours, "hour"))
	}
	if minutes > 0 {
		parts = append(parts, pluralize(minutes, "minute"))
	}
	if seconds > 0 {
		parts = append(parts, pluralize(seconds, "second"))
	}

	if len(parts) == 0 {
		// Duration was less than a second but non-zero.
		return prefix + "0 seconds"
	}

	return prefix + strings.Join(parts, " ")
}
