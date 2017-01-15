package expenses

type Expense struct {
	Id        int64  `json:"id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	ExpenseParams
}

type Expenses struct {
	Expenses []Expense `json:"expenses"`
}

type ExpenseParams struct {
	UserId     int64  `json:"user_id,omitempty"`
	PriceCents int64  `json:"price_cents,string,omitempty"`
	Date       string `json:"date,omitempty"`
	Tag        string `json:"tag,omitempty"`
	Notes      string `json:"notes,omitempty"`
}

type FilterParams struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

func NewExpense(params ExpenseParams) Expense {
	return Expense{ExpenseParams: params}
}
