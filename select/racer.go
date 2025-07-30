package selectp

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := meassureResponseTime(a)
	bDuration := meassureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func meassureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
