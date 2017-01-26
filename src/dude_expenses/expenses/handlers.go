package expenses

import (
	"dude_expenses/app"
	"net/http"
	"strconv"
)

type indexHandler struct {
	env *app.Env
}

func (h indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) app.Response {
	var params FilterParams
	w.Header().Set("Pragma", "no-cache")

	queryParams := r.URL.Query()
	params.From = queryParams.Get("from")
	params.To = queryParams.Get("to")
	params.UserId = h.env.GetUserId()

	expenses, err := NewRepository(h.env.GetDB()).GetExpenses(params)
	if err != nil {
		return app.InternalServerError()
	}

	return app.OK(expenses)
}

func Index(env *app.Env) app.Handler {
	return indexHandler{env: env}
}

type createHandler struct {
	env *app.Env
}

func (h createHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) app.Response {
	var expenseParams ExpenseParams

	err := app.ParseRequestBody(r, &expenseParams)
	defer r.Body.Close()
	if err != nil {
		return app.BadRequest()
	}
	expenseParams.UserId, err = strconv.ParseInt(h.env.GetUserId(), 10, 64)
	if err != nil {
		return app.BadRequest()
	}

	expenseValidation := NewExpenseValidation(expenseParams)
	if !expenseValidation.IsValid() {
		return app.UnprocessableEntity(expenseValidation.Errors)
	}

	repository := NewRepository(h.env.GetDB())
	expense, err := repository.CreateExpense(expenseParams)
	if err != nil {
		return app.InternalServerError()
	}

	return app.Created(expense)
}

func Create(env *app.Env) app.Handler {
	return createHandler{env: env}
}
