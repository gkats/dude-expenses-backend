package users

import (
	"dude_expenses/app"
	"net/http"
)

type createHandler struct {
	env *app.Env
}

func (h createHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) app.Response {
	var userParams UserParams

	err := app.ParseRequestBody(r, &userParams)
	defer r.Body.Close()
	if err != nil {
		return app.BadRequest()
	}

	repository := NewRepository(h.env.GetDB())
	userValidation := NewUserValidation(userParams, repository)
	if !userValidation.IsValid() {
		return app.UnprocessableEntity(userValidation.Errors)
	}

	user, err := repository.CreateUser(userParams)
	if err != nil {
		// TODO Log error
		return app.InternalServerError()
	}

	return app.Created(user)
}

func Create(env *app.Env) app.Handler {
	return createHandler{env: env}
}

type authenticateHandler struct {
	env *app.Env
}

func (h authenticateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) app.Response {
	var userParams UserParams

	err := app.ParseRequestBody(r, &userParams)
	defer r.Body.Close()
	if err != nil {
		return app.BadRequest()
	}

	authService := NewAuthService(h.env)
	token, err := authService.Authenticate(userParams)
	if err != nil {
		if _, ok := err.(*AuthInvalidEmailError); ok {
			return app.NotFound()
		}
		if _, ok := err.(*AuthInvalidPasswordError); ok {
			errors := make(map[string][]string)
			errors["password"] = append(errors["password"], "is invalid")
			return app.UnprocessableEntity(errors)
		}
		// TODO Log error
		return app.InternalServerError()
	}

	return app.OK(token)
}

func Authenticate(env *app.Env) app.Handler {
	return authenticateHandler{env: env}
}
