package utils

import "time"

var now func() time.Time = time.Now

func SetNow(fn func() time.Time) {
	now = fn
}

func Now() time.Time {
	return now()
}
