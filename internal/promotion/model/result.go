package model

type Result struct {
	Total          float64 `json:"total"`
	TotalAfterDisc float64 `json:"total_after_discount"`
	Discount       float64 `json:"discount"`
}
