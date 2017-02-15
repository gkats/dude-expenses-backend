package expenses

type ExpenseValidation struct {
	ExpenseParams
	Errors map[string][]string `json:"errors"`
}

func NewExpenseValidation(params ExpenseParams) ExpenseValidation {
	return ExpenseValidation{ExpenseParams: params}
}

func (validation *ExpenseValidation) IsValid() bool {
	validation.Errors = make(map[string][]string)

	if validation.PriceCents <= 0 {
		validation.Errors["priceCents"] = append(validation.Errors["priceCents"], "must be greater than 0")
	}

	if len(validation.Date) == 0 {
		validation.Errors["date"] = append(validation.Errors["date"], "can't be blank")
	}

	if len(validation.Tag) == 0 {
		validation.Errors["tag"] = append(validation.Errors["tag"], "can't be blank")
	}

	if validation.UserId == 0 {
		validation.Errors["userId"] = append(validation.Errors["userId"], "can't be blank")
	}

	return len(validation.Errors) == 0
}
