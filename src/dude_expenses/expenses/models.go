package expenses

type Expense struct {
	Id        int64  `json:"id,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	ExpenseParams
}

type Expenses struct {
	Expenses []Expense `json:"expenses"`
}

type ExpenseParams struct {
	UserId     int64  `json:"userId,omitempty"`
	PriceCents int64  `json:"priceCents,string,omitempty"`
	Date       string `json:"date,omitempty"`
	Tag        string `json:"tag,omitempty"`
	Notes      string `json:"notes,omitempty"`
}

type FilterParams struct {
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
	UserId string `json:"userId,omitempty"`
}

func NewExpense(params ExpenseParams) Expense {
	return Expense{ExpenseParams: params}
}

type Tags struct {
	Tags []string `json:"tags"`
}
