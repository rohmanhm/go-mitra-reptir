package response

// Products JSON schema result.
type Products struct {
	Base
	Data []ProductsData `json:"data"`
}

// ProductsData JSON schema result.
type ProductsData struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Amount int64  `json:"amount"`
	Status int64  `json:"status"`
}
