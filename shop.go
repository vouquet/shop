package shop

import (
	"time"
)

const (
	ORDER_TYPE_BUY string = "BUY"
	ORDER_TYPE_SELL string = "SELL"
)

type Shop interface {
	GetRate()   (map[string]Rate, error)
	GetPositions(string) ([]Position, error)
	OrderStreamIn(string, string, float64) error //type, symbol, size
	OrderStreamOut(Position) error
}

type Position interface {
	Id()     string
	Symbol() string
	Size()   float64
	Price()  float64
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
