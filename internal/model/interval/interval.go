package interval

type Interval string

const (
	Min1  = "1m"
	Min3  = "3m"
	Min5  = "5m"
	Min15 = "15m"
	Min30 = "30m"

	Hour1  = "1h"
	Hour2  = "2h"
	Hour4  = "4h"
	Hour6  = "6h"
	Hour8  = "8h"
	Hour12 = "12h"

	Day1   = "1d"
	Day3   = "3d"
	Week1  = "1w"
	Month1 = "1M"
)

func (i Interval) String() string {
	return string(i)
}
