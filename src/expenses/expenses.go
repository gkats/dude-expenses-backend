package expenses

type Expense struct {
	Id int64 `json:"id,omitempty"`
	ExpenseParams
}

type ExpenseParams struct {
	PriceCents int64 `json:"price_cents,string,omitempty"`
	Date string `json:"date,omitempty"`
	Tag string `json:"tag,omitempty"`
	Notes string `json:"notes,omitempty"`
}

func NewExpense(params ExpenseParams) Expense {
	return Expense{ExpenseParams: params}
}
