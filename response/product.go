package response

// Product JSON schema result.
type Product struct {
	Base
	Data ProductData `json:"data"`
}

// ProductData JSON schema result.
type ProductData struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Provider string `json:"provider"`
	Type     string `json:"type"`
	Amount   int64  `json:"amount"`
	Status   int64  `json:"status"`
}
