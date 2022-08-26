package api

import (
	"catalyst-case/database"
	"catalyst-case/module/customer"
	"catalyst-case/pkg/response"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type customerApi struct {
	customer customer.Customer
}

func newCustomerApi(db *database.DB) *customerApi {
	customerQuery := customer.NewCustomerQuery(db)
	customerCommand := customer.NewCustomerCommand(db)
	return &customerApi{
		customer: customer.NewCustomer(customerQuery, customerCommand),
	}
}

func (c *customerApi) getCustomerById(w http.ResponseWriter, r *http.Request, customerId int) {
	res, err := c.customer.GetCustomerById(r.Context(), customerId)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	response.ResponseWithJSON(w, http.StatusOK, res)
}

func (c *customerApi) getCustomerByEmail(w http.ResponseWriter, r *http.Request, customerEmail string) {
	res, err := c.customer.GetCustomerByEmail(r.Context(), customerEmail)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	response.ResponseWithJSON(w, http.StatusOK, res)
}

func (c *customerApi) postNewCustomer(w http.ResponseWriter, r *http.Request) {
	var model *customer.NewCustomerDto
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
	err = c.customer.AddNewCustomer(r.Context(), *model)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	response.ResponseWithNoContent(w)
}
