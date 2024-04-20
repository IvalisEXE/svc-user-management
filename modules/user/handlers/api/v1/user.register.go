package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"svc-user-management/lib/response"
)

func (h *Handler) RegisterUser(c echo.Context) error {
	/* init tracer */

	ctx := c.Request().Context()
	bodyReq := new(UserRegisterBodyReq)
	if err := c.Bind(bodyReq); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.NewBadRequestResponse(err.Error()),
		)
	}

	if err := bodyReq.RegiterUserValidation(); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.NewBadRequestResponse(err.Error()),
		)
	}

	if err := bodyReq.NormalizeMSISDN(); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.NewBadRequestResponse(err.Error()),
		)
	}

	userExist, err := h.userUsecases.CheckUserExistByPhoneNo(ctx, bodyReq.PhoneNo)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			err,
		)
	}

	if userExist {
		return c.JSON(
			http.StatusUnprocessableEntity,
			response.NewUnprocessableEntity("user is already exist"),
		)
	}

	userID, err := h.userUsecases.CreateUser(ctx, bodyReq.ParseToRegisterUserEntity())
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			response.NewInternalServerError("internal server error"),
		)
	}

	bodyRes := new(UserRegisterBodyRes)
	bodyRes.ToRegisterUserResponse(userID)
	return c.JSON(
		http.StatusCreated,
		response.NewSuccessResponseCreated(
			response.StatusCodeCreated,
			"register user successfully...",
			bodyRes,
		),
	)
}
