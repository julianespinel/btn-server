package alert

import "time"

type Alert struct {
	Id        float64   `json:"id"`
	Serial    string    `json:"serial"`
	ElderId   string    `json:"elderId"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Date      time.Time `json:"date"`
}

type SendingResult struct {
	Id            float64 `json:"id"`
	AlertId       float64 `json:"alertId"`
	RelativeId    string  `json:"relativeId"`
	WasSuccessful bool    `json:"wasSuccessful"`
	ResultMessage string  `json: resultMessage`
}
