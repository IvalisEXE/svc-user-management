package v1

import (
	"net/http"
	"strconv"
	"svc-user-management/lib/response"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetProfileUser(c echo.Context) error {
	ctx := c.Request().Context()
	userID := c.Get("user_id").(string)
	numUserID, _ := strconv.ParseInt(userID, 10, 64)
	user, err := h.userUsecases.GetUserByID(ctx, numUserID)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			response.NewInternalServerError("internal server error"),
		)
	}

	if user == nil {
		return c.JSON(
			http.StatusForbidden,
			response.NewForbiddenResponse(),
		)
	}

	userProfile := new(UserProfileBodyRes)
	userProfile.UserProfileUserBodyRes.PhoneNo = user.PhoneNo
	userProfile.UserProfileUserBodyRes.Fullname = user.FullName

	return c.JSON(
		http.StatusOK,
		response.NewSuccessResponseStatusOK(
			response.StatusCodeCreated,
			"get user successfully...",
			userProfile,
		),
	)
}

func (h *Handler) UpdateProfileUser(c echo.Context) error {
	ctx := c.Request().Context()

	bodyReq := new(UpdateUserProfileBodyReq)
	if err := c.Bind(bodyReq); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.NewBadRequestResponse(err.Error()),
		)
	}

	skipPhoneNo := bodyReq.PhoneNo == ""
	if !skipPhoneNo {
		if err := bodyReq.UpdateUserProfileValidation(); err != nil {
			return c.JSON(
				http.StatusBadRequest,
				response.NewBadRequestResponse(err.Error()),
			)
		}

		userExist, err := h.userUsecases.CheckUserExistByPhoneNo(ctx, bodyReq.PhoneNo)
		if err != nil {
			return c.JSON(
				http.StatusBadRequest,
				response.NewBadRequestResponse(err.Error()),
			)
		}

		if userExist {
			return c.JSON(
				http.StatusConflict,
				response.NewConflict("phone  number is already exist"),
			)
		}
	}

	userID := c.Get("user_id").(string)
	numUserID, _ := strconv.ParseInt(userID, 10, 64)
	if err := h.userUsecases.UpdateUserProfileByID(ctx, numUserID, bodyReq.PhoneNo, bodyReq.FullName, skipPhoneNo); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.NewBadRequestResponse(err.Error()),
		)
	}

	return c.JSON(
		http.StatusOK,
		response.NewSuccessResponseStatusOK(
			response.StatusCodeCreated,
			"update user successfully...",
			"ok",
		),
	)
}
