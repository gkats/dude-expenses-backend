package expenses

import (
	"encoding/json"
	"net/http"
	"app"
)

type JsonError struct {
	Message string `json:"message"`
}

type unprocessableEntityResponse struct {
	Message string `json:"message"`
	Errors map[string][]string `json:"errors"`
}

func NewUnprocessableEntityResponse(errors map[string][]string) (unprocessableEntityResponse) {
	return unprocessableEntityResponse{Message: "Resource invalid", Errors: errors }
}

func Create(env *app.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var expenseParams ExpenseParams

		w.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(w)

		err := parseRequestBody(r, &expenseParams)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(&JsonError{Message: "Bad request"})
			return
		}

		expenseValidation := NewExpenseValidation(expenseParams)
		if !expenseValidation.IsValid() {
			w.WriteHeader(http.StatusUnprocessableEntity)
			response := NewUnprocessableEntityResponse(expenseValidation.Errors)
			encoder.Encode(&response)
			return
		}

		repository := NewRepository(env.GetDB())
		expense, err := repository.CreateExpense(expenseParams)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(&JsonError{Message: "Something went wrong"})
			return
		}

		w.WriteHeader(http.StatusCreated)
		encoder.Encode(&expense)
	})
}

func parseRequestBody(r *http.Request, params *ExpenseParams) error {
	return json.NewDecoder(r.Body).Decode(&params)
}