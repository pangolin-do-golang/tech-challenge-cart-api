package controller

import (
	"net/http"

	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/core/cart"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/errutil"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartController struct {
	AbstractController
	service cart.IService
}

func NewCartController(cart cart.IService) *CartController {
	return &CartController{
		service: cart,
	}
}

type AddProductPayload struct {
	ClientID  uuid.UUID `json:"client_id" binding:"required" format:"uuid"`
	ProductID uuid.UUID `json:"product_id" binding:"required" format:"uuid"`
	Quantity  int       `json:"quantity" binding:"required,min=1" example:"1"`
	Comments  string    `json:"comments"`
}

// AddProduct adds a Product to Customer's Cart
// @Description Adds a Product to Customer's Cart
// @Tags Cart
// @Param payload body controller.AddProductPayload true "AddProductPayload"
// @Accept json
// @Produce json
// @Success 200
// @Failure 500 {object} map[string]any "{\"error\": \"Internal Server Error\"}"
// @Router /cart/add-product [post]
func (ctrl CartController) AddProduct(c *gin.Context) {
	payload := &AddProductPayload{}
	err := c.BindJSON(payload)
	if err != nil {
		ctrl.Error(c, errutil.NewInputError(err))
		return
	}

	err = ctrl.service.AddProduct(c.Request.Context(), payload.ClientID, &cart.Product{
		ProductID: payload.ProductID,
		Quantity:  payload.Quantity,
		Comments:  payload.Comments,
	})

	if err != nil {
		ctrl.Error(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type EditProductPayload struct {
	ClientID  uuid.UUID `json:"client_id" binding:"required" format:"uuid"`
	ProductID uuid.UUID `json:"product_id" binding:"required" format:"uuid"`
	Quantity  int       `json:"quantity" binding:"required" example:"2"`
	Comments  string    `json:"comments"`
}

// EditProduct edits a Product from Cart by ID
// @Description Edits a Product from Customer's Cart
// @Tags Cart
// @Param payload body controller.EditProductPayload true "EditProductPayload"
// @Accept json
// @Produce json
// @Success 200
// @Failure 500 {object} map[string]any "{\"error\": \"Internal Server Error\"}"
// @Router /cart/edit-product [post]
func (ctrl CartController) EditProduct(c *gin.Context) {
	payload := &EditProductPayload{}
	err := c.BindJSON(payload)
	if err != nil {
		ctrl.Error(c, errutil.NewInputError(err))
		return
	}

	err = ctrl.service.EditProduct(c.Request.Context(), payload.ClientID, &cart.Product{
		ProductID: payload.ProductID,
		Quantity:  payload.Quantity,
		Comments:  payload.Comments,
	})
	if err != nil {
		ctrl.Error(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type RemoveProductPayload struct {
	ClientID  uuid.UUID `json:"client_id" binding:"required" format:"uuid"`
	ProductID uuid.UUID `json:"product_id" binding:"required" format:"uuid"`
}

// RemoveProduct removes a Product from Customer's Cart
// @Description Removes a Product from Customer's Cart
// @Tags Cart
// @Param payload body controller.RemoveProductPayload true "RemoveProductPayload"
// @Accept json
// @Produce json
// @Success 200
// @Failure 500 "Internal Server Error"
// @Router /cart/remove-product [post]
func (ctrl CartController) RemoveProduct(c *gin.Context) {
	payload := &RemoveProductPayload{}
	err := c.BindJSON(payload)
	if err != nil {
		ctrl.Error(c, errutil.NewInputError(err))
		return
	}

	err = ctrl.service.RemoveProduct(c.Request.Context(), payload.ClientID, payload.ProductID)
	if err != nil {
		ctrl.Error(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type GetCartPayload struct {
	ClientID uuid.UUID `json:"client_id" binding:"required" format:"uuid"`
}

// Overview get the full cart and its products
// @Tags Cart
// @Param payload body controller.GetCartPayload true "GetCartPayload"
// @Accept json
// @Produce json
// @Success 200
// @Failure 500 "Internal Server Error"
// @Router /cart/overview [post]
func (ctrl CartController) Overview(c *gin.Context) {
	payload := &GetCartPayload{}
	err := c.BindJSON(payload)
	if err != nil {
		ctrl.Error(c, errutil.NewInputError(err))
		return
	}

	fullcart, err := ctrl.service.GetFullCart(payload.ClientID)
	if err != nil {
		ctrl.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, fullcart)
}

type CleanupPayload struct {
	ClientID uuid.UUID `json:"client_id" binding:"required" format:"uuid"`
}

func (ctrl CartController) Cleanup(c *gin.Context) {
	payload := &CleanupPayload{}

	err := c.BindJSON(payload)

	if err != nil {
		ctrl.Error(c, errutil.NewInputError(err))
		return
	}

	err = ctrl.service.Cleanup(payload.ClientID)

	if err != nil {
		ctrl.Error(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type LoadCardPayload struct {
	ClientID uuid.UUID `json:"client_id" binding:"required" format:"uuid"`
}

func (ctrl CartController) LoadCart(c *gin.Context) {
	payload := &LoadCardPayload{}

	err := c.BindJSON(payload)

	if err != nil {
		ctrl.Error(c, errutil.NewInputError(err))
		return
	}

	cart, err := ctrl.service.LoadCart(payload.ClientID)

	if err != nil {
		ctrl.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, cart)
}
