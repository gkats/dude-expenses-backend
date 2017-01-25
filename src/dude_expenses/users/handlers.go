package users

import (
	"dude_expenses/app"
	"dude_expenses/app/handler"
	"net/http"
)

func Create(env *app.Env) handler.AppHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (handler.HandlerResponse, handler.Error) {
		var userParams UserParams
		var response handler.HandlerResponse

		err := handler.ParseRequestBody(r, &userParams)
		defer r.Body.Close()
		if err != nil {
			// TODO Log error??
			return response, handler.BadRequest()
		}

		repository := NewRepository(env.GetDB())
		userValidation := NewUserValidation(userParams, repository)
		if !userValidation.IsValid() {
			return response, handler.UnprocessableEntity(userValidation.Errors)
		}

		user, err := repository.CreateUser(userParams)
		if err != nil {
			// TODO Log error
			return response, handler.InternalServerError()
		}

		response = handler.NewHandlerResponse(http.StatusCreated, user)
		return response, nil
	}
}

func Authenticate(env *app.Env) handler.AppHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (handler.HandlerResponse, handler.Error) {
		var userParams UserParams
		var response handler.HandlerResponse

		err := handler.ParseRequestBody(r, &userParams)
		defer r.Body.Close()
		if err != nil {
			return response, handler.BadRequest()
		}

		authService := NewAuthService(env)
		token, err := authService.Authenticate(userParams)
		if err != nil {
			if _, ok := err.(*AuthInvalidEmailError); ok {
				return response, handler.NotFound()
			}
			if _, ok := err.(*AuthInvalidPasswordError); ok {
				errors := make(map[string][]string)
				errors["password"] = append(errors["password"], "is invalid")
				return response, handler.UnprocessableEntity(errors)
			}
			// TODO Log error
			return response, handler.InternalServerError()
		}

		response = handler.NewHandlerResponse(http.StatusOK, token)
		return response, nil
	}
}
