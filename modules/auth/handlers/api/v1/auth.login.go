package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"svc-user-management/lib/custerr"
	"svc-user-management/lib/response"
)

func (h *Handler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	bodyReq := c.Get("credential").(Credential)
	if err := bodyReq.LengthValidation(); err != nil {
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

	token, userID, err := h.authUsecases.Login(ctx, bodyReq.PhoneNo, bodyReq.Password)
	if err != nil {
		if err.Error() == custerr.ErrNoRows.Error() {
			return c.JSON(
				http.StatusUnprocessableEntity,
				response.NewUnprocessableEntity(custerr.ErrUserNotFound.Error()),
			)

		}
		if err.Error() == custerr.ErrInvalidCredentials.Error() {
			return c.JSON(
				http.StatusUnauthorized,
				response.NewUnauthorizedWithMessage(custerr.ErrUnauthorized.Error()),
			)
		}
		return c.JSON(
			http.StatusInternalServerError,
			response.NewInternalServerError("internal server error"),
		)
	}

	bodyRes := new(LoginBodyRes)
	bodyRes.User.UserID = strconv.FormatInt(userID, 10)
	bodyRes.Token = token

	return c.JSON(
		http.StatusOK,
		response.NewSuccessResponseStatusOK(
			response.StatusCodeSuccess,
			"login user successfully...",
			bodyRes,
		),
	)
}
