package lib

import "time"

func TimePtr(time time.Time) *time.Time {
	return &time
}

func Float64Ptr(data float64) *float64 {
	return &data
}

func BoolPtr(data bool) *bool {
	return &data
}
