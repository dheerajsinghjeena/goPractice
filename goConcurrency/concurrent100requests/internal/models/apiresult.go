package models

import "time"

type ResultStruct struct {
	ResultID     int
	URL          string
	StatusCode   int
	Body         any
	ApiError     error
	ResponseTime time.Time
}
