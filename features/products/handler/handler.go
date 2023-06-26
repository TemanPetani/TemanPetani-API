package handler

import (
	"alta/temanpetani/features/products"
	"alta/temanpetani/utils/helpers"
	"alta/temanpetani/utils/middlewares"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService products.ProductServiceInterface
}

func (handler *UserHandler) PostProductHandler(c echo.Context) error {
	var payload products.Core
	if errBind := c.Bind(&payload); errBind != nil {
		return helpers.StatusBadRequestResponse(c, "error bind payload: " + errBind.Error())
	}
	file, err := c.FormFile("image");
	if err != nil {
		return helpers.StatusBadRequestResponse(c, "error get file image: " + err.Error())
	}
	payload.Image = file

	userId, _, errExtractUserId := middlewares.ExtractToken(c)
	if errExtractUserId != nil {
		return helpers.StatusAuthorizationErrorResponse(c, "error get user id: " + errExtractUserId.Error())
	}
	payload.UserID = uint(userId)

	productId, err := handler.userService.AddProduct(payload)
	if err != nil {
		if strings.Contains(err.Error(), "validator") {
			return helpers.StatusBadRequestResponse(c, err.Error())
		} else {
			return helpers.StatusInternalServerError(c, err.Error())
		}
	}
	if productId != "" {
		return helpers.StatusCreated(c, "Berhasil menambahkan product", map[string]any{
			"productId": productId, 
		})
	}
	return nil
}

func (handler *UserHandler) PutImageProductHandler(c echo.Context) error {
	var payload products.CoreProductImageRequest
	productId := c.Param("id")
	file, err := c.FormFile("image");
	if err != nil {
		return helpers.StatusBadRequestResponse(c, "error get file image: " + err.Error())
	}
	payload.Image = file

	imageUrl, errUpdate := handler.userService.UpdateImage(productId, payload)
	if errUpdate != nil {
		if strings.Contains(errUpdate.Error(), "validator") {
			return helpers.StatusBadRequestResponse(c, errUpdate.Error())
		} else {
			return helpers.StatusInternalServerError(c, errUpdate.Error())
		}
	}

	if imageUrl != "" {
		return helpers.StatusOKWithData(c, "Berhasil mengupdate image product", map[string]any{
			"imageUrl": imageUrl,
		})
	}
	return nil
}

func New(userService products.ProductServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}