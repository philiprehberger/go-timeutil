# timeutil

A zero-dependency Go library for time manipulation helpers. Provides boundary calculations, date comparisons, and human-readable time formatting.

## Install

```sh
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
| `DaysBetween(a, b)` | Absolute number of calendar days between two times |
| `IsWeekend(t)` | True if Saturday or Sunday |
| `IsBusinessDay(t)` | True if Monday through Friday |
| `NextBusinessDay(t)` | Start of next weekday after t |
| `Humanize(t)` | Relative time string ("3 hours ago", "in 5 minutes") |
| `HumanizeDuration(d)` | Human-readable duration ("2 hours 30 minutes") |

All boundary functions preserve the original `time.Location`.

## License

MIT
