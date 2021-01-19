package model

type RawPayload struct {
	Cart []Payload `json:"cart"`
}

type Payload struct {
	SKU   string  `json:"sku"`
	Price float64 `json:"price"`
	Qty   int     `json:"qty"`
}
