package api

import (
	"catalyst-case/database"
	"catalyst-case/module/order"
	"catalyst-case/module/product"
	"catalyst-case/pkg/response"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type orderApi struct {
	order order.Order
}

func newOrderApi(db *database.DB) *orderApi {
	orderCommand := order.NewOrderCommand(db)
	orderQuery := order.NewOrderQuery(db)
	productQuery := product.NewProductQuery(db)
	return &orderApi{
		order: order.NewOrder(orderCommand, productQuery, orderQuery),
	}
}

func (o *orderApi) postNewOrder(w http.ResponseWriter, r *http.Request) {
	var model *order.OrderEntry
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	validate := validator.New()
	err = validate.Struct(model)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	res, err := o.order.AddNewOrder(r.Context(), *model)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	response.ResponseWithJSON(w, http.StatusOK, res)
}

func (o *orderApi) getOrderById(w http.ResponseWriter, r *http.Request, orderId int) {
	res, err := o.order.GetOrder(r.Context(), orderId)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	response.ResponseWithJSON(w, http.StatusOK, res)
}

func (o *orderApi) getOrders(w http.ResponseWriter, r *http.Request) {
	res, err := o.order.GetOrders(r.Context())
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	response.ResponseWithJSON(w, http.StatusOK, res)
}
