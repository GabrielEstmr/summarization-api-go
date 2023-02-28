package main_gateways_http_controllers

import (
	"encoding/json"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"io"
	main_domains "mpindicator/main/domains"
	main_gateways_http_controllers_requests "mpindicator/main/gateways/http/controllers/request"

	main_usecases "mpindicator/main/usecases"
	main_utils "mpindicator/main/utils"
	"net/http"
)

const IDX_TRACING_FIND_USER_CONTROLLER = "find-user-controller"

type UserController struct {
	useCaseCreateUser main_usecases.CreateUser
}

func NewUserController(useCaseCreateUser *main_usecases.CreateUser) *UserController {
	return &UserController{useCaseCreateUser: *useCaseCreateUser}
}

func (thisController *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	ctx := baggage.ContextWithoutBaggage(r.Context())
	tr := otel.GetTracerProvider().Tracer(main_domains.APP_INDICATOR_TYPE_FIND_USER.GetDescription())
	ctx, controller := tr.Start(ctx, IDX_TRACING_FIND_USER_CONTROLLER)
	defer controller.End()

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		main_utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var userRequest main_gateways_http_controllers_requests.UserRequest
	if err = json.Unmarshal(requestBody, &userRequest); err != nil {
		main_utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	errValidate := userRequest.Validate()
	if errValidate != nil {
		main_utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := userRequest.ToDomain()

	id, err := thisController.useCaseCreateUser.Execute(ctx, user)
	if err != nil {
		main_utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	main_utils.JSON(w, http.StatusAccepted, id)
}
