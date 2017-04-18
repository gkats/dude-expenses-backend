package expenses

import (
	"dude_expenses/app"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type indexHandler struct {
	env *app.Env
}

func (h indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) app.Response {
	var params FilterParams

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

type tagsHandler struct {
	env *app.Env
}

func (h tagsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) app.Response {
	tags, err := NewRepository(h.env.GetDB()).GetTags(h.env.GetUserId())
	if err != nil {
		return app.InternalServerError()
	}

	return app.OK(tags)
}

func GetTags(env *app.Env) app.Handler {
	return tagsHandler{env: env}
}

type showHandler struct {
	env *app.Env
}

func (h showHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) app.Response {
	vars := mux.Vars(r)
	expense, err := NewRepository(h.env.GetDB()).GetExpense(vars["id"])
	if err != nil {
		return app.InternalServerError()
	}
	if expense == nil {
		return app.NotFound()
	}

	return app.OK(expense)
}

func GetExpense(env *app.Env) app.Handler {
	return showHandler{env: env}
}
