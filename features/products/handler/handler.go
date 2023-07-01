package handler

import (
	"alta/temanpetani/features/products"
	"alta/temanpetani/utils/helpers"
	"alta/temanpetani/utils/middlewares"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	productService products.ProductServiceInterface
}

func (handler *UserHandler) PostProductHandler(c echo.Context) error {
	var payload products.Core
	if errBind := c.Bind(&payload); errBind != nil {
		return helpers.StatusBadRequestResponse(c, "error bind payload: "+errBind.Error())
	}
	file, err := c.FormFile("image")
	if err != nil {
		return helpers.StatusBadRequestResponse(c, "error get file image: "+err.Error())
	}
	payload.Image = file

	userId, _, errExtractUserId := middlewares.ExtractToken(c)
	if errExtractUserId != nil {
		return helpers.StatusAuthorizationErrorResponse(c, "error get user id: "+errExtractUserId.Error())
	}
	payload.UserID = userId

	productId, err := handler.productService.AddProduct(payload)
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
	file, err := c.FormFile("image")
	if err != nil {
		return helpers.StatusBadRequestResponse(c, "error get file image: "+err.Error())
	}
	payload.Image = file

	imageUrl, errUpdate := handler.productService.UpdateImage(productId, payload)
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

func (handler *UserHandler) GetAllProductsHandler(c echo.Context) error {
	querys := map[string]any{}
	role := c.QueryParam("role")
	if role != "" {
		querys["role"] = role
	}

	products, err := handler.productService.GetAllProducts(querys)
	if err != nil {
		return helpers.StatusInternalServerError(c, err.Error())
	}
	return helpers.StatusOKWithData(c, "Berhasil mendapatkan sejumlah produk", map[string]any{
		"products": products,
	})
}

func (handler *UserHandler) GetProductsByUserIdHandler(c echo.Context) error {
	userId, _, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		return helpers.StatusBadRequestResponse(c, errExtract.Error())
	}

	products, err := handler.productService.GetProductsByUserId(userId)
	if err != nil {
		return helpers.StatusInternalServerError(c, err.Error())
	}
	return helpers.StatusOKWithData(c, "Berhasil mendapatkan sejumlah produk", map[string]any{
		"products": products,
	})
}

func New(productService products.ProductServiceInterface) *UserHandler {
	return &UserHandler{
		productService: productService,
	}
}
