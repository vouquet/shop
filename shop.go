package shop

import (
	"time"
)

type Shop interface {
	GetRate()   (map[string]Rate, error)
}

type Rate interface {
	Ask()       float64
	Bid()       float64
	High()      float64
	Last()      float64
	Low()       float64
	Symbol()    string
	Time()      time.Time
	Volume()    float64
}
