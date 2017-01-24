package expenses

import (
	"app"
	"app/handler"
	"auth"
	"fmt"
	"net/http"
)

func Create(env *app.Env) handler.AppHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (handler.HandlerResponse, handler.Error) {
		var apiResponse handler.HandlerResponse
		var expenseParams ExpenseParams

		if err := auth.Authenticate(env, r); err != nil {
			return apiResponse, handler.Unauthorized()
		}

		err := handler.ParseRequestBody(r, &expenseParams)
		defer r.Body.Close()
		if err != nil {
			// TODO Log err??
			return apiResponse, handler.BadRequest()
		}

		expenseValidation := NewExpenseValidation(expenseParams)
		if !expenseValidation.IsValid() {
			return apiResponse, handler.UnprocessableEntity(expenseValidation.Errors)
		}

		repository := NewRepository(env.GetDB())
		expense, err := repository.CreateExpense(expenseParams)
		if err != nil {
			// TODO Log error!!
			return apiResponse, handler.InternalServerError()
		}

		return handler.NewHandlerResponse(http.StatusCreated, expense), nil
	}
}

func Index(env *app.Env) handler.AppHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (handler.HandlerResponse, handler.Error) {
		var apiResponse handler.HandlerResponse
		var params FilterParams

		if err := auth.Authenticate(env, r); err != nil {
			return apiResponse, handler.Unauthorized()
		}

		w.Header().Set("Pragma", "no-cache")

		queryParams := r.URL.Query()
		fmt.Println("UserId: " + env.GetUserId())
		params.From = queryParams.Get("from")
		params.To = queryParams.Get("to")

		expenses, err := NewRepository(env.GetDB()).GetExpenses(params)
		if err != nil {
			// TODO Log error!
			return apiResponse, handler.InternalServerError()
		}

		return handler.NewHandlerResponse(http.StatusOK, expenses), nil
	}
}
