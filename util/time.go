package util

import "time"

// GetTimeNow はTZを固定する為のラッパーです。
func GetTimeNow() time.Time {
	return time.Now().In(time.UTC)
}
