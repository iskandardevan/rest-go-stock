package product_types

import (
	"net/http"
	"rest-go-stock/businesses/product_types"
	"rest-go-stock/controllers"
	"rest-go-stock/controllers/product_types/request"
	"rest-go-stock/controllers/product_types/response"

	"github.com/labstack/echo/v4"
)

type ProductTypeController struct {
	productTypeUseCase product_types.ProductTypeUsecaseInterface
}

func NewProductTypeController(ProductTypeUseCase product_types.ProductTypeUsecaseInterface) *ProductTypeController {
	return &ProductTypeController{productTypeUseCase: ProductTypeUseCase}
}

func (ctrl *ProductTypeController) AddProductType(c echo.Context) error {
	req := request.ProductTypeRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	var data product_types.Domain
	data, err = ctrl.productTypeUseCase.AddProductType(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainProductType(data))
}


// func (ctrl *ProductTypeController) GetByID(ctx echo.Context) error {
// 	req := ctx.Request().Context()
// 	id := ctx.Param("id")
// 	Convint, errConvint := strconv.Atoi(id)
// 	if errConvint != nil {
// 		return controllers.NewErrorResponse(ctx, http.StatusBadRequest, errConvint)
// 	}
// 	data, err := ProductTypeController.productTypeUseCase.GetByID(uint(Convint), req)
// 	if err != nil {
// 		return controllers.NewErrorResponse(ctx, http.StatusInternalServerError, err)
// 	}
// 	return controllers.NewSuccesResponse(ctx, response.FromDomainProductType(data))
// }



