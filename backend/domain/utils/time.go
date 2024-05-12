package utils

import "time"

var now func() time.Time = time.Now

func Now() time.Time {
	return now()
}
