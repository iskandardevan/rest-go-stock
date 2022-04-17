package products

import (
	"net/http"
	"rest-go-stock/businesses/products"
	"rest-go-stock/controllers"
	"rest-go-stock/controllers/products/request"
	"rest-go-stock/controllers/products/response"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productUseCase products.ProductUsecaseInterface
}

func NewProductController(ProductUseCase products.ProductUsecaseInterface) *ProductController {
	return &ProductController{productUseCase: ProductUseCase}
}

func (ctrl *ProductController) Add(c echo.Context) error {
	req := request.ProductRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	var data products.Domain
	data, err = ctrl.productUseCase.Add(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomainProduct(data))
}
