package model

import "time"

// CardInfo represents the card's code and price
type CardInfo struct {
	ImageURL    string
	Code        string
	CardName    string
	HrefLink    string
	YenPrice    string
	UpdatedTime time.Time
}
