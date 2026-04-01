# go-timeutil

[![CI](https://github.com/philiprehberger/go-timeutil/actions/workflows/ci.yml/badge.svg)](https://github.com/philiprehberger/go-timeutil/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/philiprehberger/go-timeutil.svg)](https://pkg.go.dev/github.com/philiprehberger/go-timeutil)
[![Last updated](https://img.shields.io/github/last-commit/philiprehberger/go-timeutil)](https://github.com/philiprehberger/go-timeutil/commits/main)

Time manipulation helpers for Go — boundaries, comparisons, and human-readable formatting

## Installation

```bash
go get github.com/philiprehberger/go-timeutil
```

## Usage

```go
package main

import (
    "fmt"
    "time"

    "github.com/philiprehberger/go-timeutil"
)

func main() {
    now := time.Now()
    later := now.AddDate(0, 0, 14)

    // Boundary helpers
    fmt.Println(timeutil.StartOfDay(now))   // Today at 00:00:00
    fmt.Println(timeutil.EndOfDay(now))     // Today at 23:59:59.999999999
    fmt.Println(timeutil.StartOfWeek(now))  // Monday 00:00:00
    fmt.Println(timeutil.StartOfMonth(now)) // 1st of month 00:00:00

    // Date helpers
    fmt.Println(timeutil.DaysBetween(now, now.AddDate(0, 0, 10))) // 10
    fmt.Println(timeutil.IsWeekend(now))                          // true/false
    fmt.Println(timeutil.IsBusinessDay(now))                      // true/false
    fmt.Println(timeutil.NextBusinessDay(now))                    // Next Mon-Fri

    // Quarter boundaries
    fmt.Println(timeutil.StartOfQuarter(now)) // First day of quarter at 00:00:00
    fmt.Println(timeutil.EndOfQuarter(now))   // Last day of quarter at 23:59:59.999999999

    // Year boundaries
    fmt.Println(timeutil.StartOfYear(now)) // Jan 1 at 00:00:00
    fmt.Println(timeutil.EndOfYear(now))   // Dec 31 at 23:59:59.999999999

    // Business day arithmetic
    fmt.Println(timeutil.PreviousBusinessDay(now))        // Most recent weekday before now
    fmt.Println(timeutil.AddBusinessDays(now, 5))         // 5 business days from now
    fmt.Println(timeutil.AddBusinessDays(now, -3))        // 3 business days ago
    fmt.Println(timeutil.BusinessDaysBetween(now, later)) // Weekdays between dates

    // Human-readable output
    past := now.Add(-3 * time.Hour)
    fmt.Println(timeutil.Humanize(past)) // "3 hours ago"

    d := 2*time.Hour + 30*time.Minute
    fmt.Println(timeutil.HumanizeDuration(d)) // "2 hours 30 minutes"
}
```

## API

| Function | Description |
|---|---|
| `StartOfDay(t)` | Midnight of the given day |
| `EndOfDay(t)` | 23:59:59.999999999 of the given day |
| `StartOfWeek(t)` | Monday 00:00:00 of the week |
| `EndOfWeek(t)` | Sunday 23:59:59.999999999 of the week |
| `StartOfMonth(t)` | 1st of month at 00:00:00 |
| `EndOfMonth(t)` | Last day of month at 23:59:59.999999999 |
| `StartOfQuarter(t)` | First day of quarter at 00:00:00 |
| `EndOfQuarter(t)` | Last day of quarter at 23:59:59.999999999 |
| `StartOfYear(t)` | January 1 at 00:00:00 |
| `EndOfYear(t)` | December 31 at 23:59:59.999999999 |
| `DaysBetween(a, b)` | Absolute number of calendar days between two times |
| `IsWeekend(t)` | True if Saturday or Sunday |
| `IsBusinessDay(t)` | True if Monday through Friday |
| `NextBusinessDay(t)` | Start of next weekday after t |
| `PreviousBusinessDay(t)` | Start of most recent weekday before t |
| `AddBusinessDays(t, n)` | Add/subtract n business days (skips weekends) |
| `BusinessDaysBetween(a, b)` | Count of weekdays between two dates |
| `Humanize(t)` | Relative time string ("3 hours ago", "in 5 minutes") |
| `HumanizeDuration(d)` | Human-readable duration ("2 hours 30 minutes") |

All boundary functions preserve the original `time.Location`.

## Development

```bash
go test ./...
go vet ./...
```

## Support

If you find this project useful:

⭐ [Star the repo](https://github.com/philiprehberger/go-timeutil)

🐛 [Report issues](https://github.com/philiprehberger/go-timeutil/issues?q=is%3Aissue+is%3Aopen+label%3Abug)

💡 [Suggest features](https://github.com/philiprehberger/go-timeutil/issues?q=is%3Aissue+is%3Aopen+label%3Aenhancement)

❤️ [Sponsor development](https://github.com/sponsors/philiprehberger)

🌐 [All Open Source Projects](https://philiprehberger.com/open-source-packages)

💻 [GitHub Profile](https://github.com/philiprehberger)

🔗 [LinkedIn Profile](https://www.linkedin.com/in/philiprehberger)

## License

[MIT](LICENSE)
