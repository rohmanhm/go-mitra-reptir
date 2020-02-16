package response

// BuyProduct JSON schema result.
type BuyProduct struct {
	Base
	Data BuyProductData `json:"data"`
}

// BuyProductData JSON schema result.
type BuyProductData struct {
	ID       string `json:"id"`
	Ref      string `json:"ref"`
	Type     string `json:"type"`
	Code     string `json:"code"`
	Dest     string `json:"dest"`
	Amount   int64  `json:"amount"`
	Status   string `json:"status"`
	Response string `json:"response"`
}
