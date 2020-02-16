package response

// Base result
type Base struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Errors  Errors      `json:"errors"`
}

// Errors result
type Errors struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
