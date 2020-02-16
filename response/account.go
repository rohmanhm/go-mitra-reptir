package response

// Account result data schema.
type Account struct {
	Base
	Data AccountData `json:"data"`
}

// AccountData result JSON schema.
type AccountData struct {
	UserID      int64                  `json:"userID"`
	Name        string                 `json:"name"`
	Email       string                 `json:"email"`
	Store       string                 `json:"store"`
	Balance     int64                  `json:"balance"`
	Transaction AccountDataTransaction `json:"transaction"`
}

// AccountDataTransaction result.
type AccountDataTransaction struct {
	Pending  int64 `json:"pending"`
	Complete int64 `json:"complete"`
}
