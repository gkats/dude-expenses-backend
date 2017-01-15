package users

import (
	"app"
	"app/handler"
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

		return handler.NewHandlerResponse(http.StatusCreated, user), nil
	}
}
