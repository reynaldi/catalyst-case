package api

import (
	"catalyst-case/config"
	"catalyst-case/database"
	"catalyst-case/pkg/response"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Api struct {
	cfg *config.Config
	db  *database.DB
}

type Router interface {
	CreateRouter() *http.ServeMux
}

func NewApi(cfg *config.Config, db *database.DB) *Api {
	return &Api{
		cfg: cfg,
		db:  db,
	}
}

func (api *Api) CreateRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/brands", api.brandHandler)
	mux.HandleFunc("/customers", api.customerHandler)
	mux.HandleFunc("/orders", api.orderHandler)
	mux.HandleFunc("/products", api.productHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}

func (api *Api) brandHandler(w http.ResponseWriter, r *http.Request) {
	brandApi := newBrandApi(api.db)
	switch r.Method {
	case http.MethodPost:
		brandApi.postBrand(w, r)
		return
	case http.MethodGet:
		if !r.URL.Query().Has("id") {
			brandApi.getAllBrands(w, r)
			return
		}
		brandId, e := strconv.Atoi(r.URL.Query().Get("id"))
		if e != nil {
			log.Println(e)
			response.ResponseError(w, http.StatusBadRequest, e)
			return
		}
		brandApi.getBrandById(w, r, brandId)
		return
	}
}

func (api *Api) customerHandler(w http.ResponseWriter, r *http.Request) {
	customerApi := newCustomerApi(api.db)
	switch r.Method {
	case http.MethodGet:
		if r.URL.Query().Has("id") {
			customerId, e := strconv.Atoi(r.URL.Query().Get("id"))
			if e != nil {
				log.Println(e)
				response.ResponseError(w, http.StatusBadRequest, e)
				return
			}
			customerApi.getCustomerById(w, r, customerId)
			return
		}
		if r.URL.Query().Has("email") {
			customerApi.getCustomerByEmail(w, r, r.URL.Query().Get("email"))
			return
		}
	case http.MethodPost:
		customerApi.postNewCustomer(w, r)
		return
	}
}

func (api *Api) orderHandler(w http.ResponseWriter, r *http.Request) {
	orderApi := newOrderApi(api.db)
	switch r.Method {
	case http.MethodGet:
		if r.URL.Query().Has("id") {
			orderId, e := strconv.Atoi(r.URL.Query().Get("id"))
			if e != nil {
				log.Println(e)
				response.ResponseError(w, http.StatusBadRequest, e)
				return
			}
			orderApi.getOrderById(w, r, orderId)
			return
		}
		orderApi.getOrders(w, r)
		return
	case http.MethodPost:
		orderApi.postNewOrder(w, r)
		return
	}
}

func (api *Api) productHandler(w http.ResponseWriter, r *http.Request) {
	productApi := newProductApi(api.db)
	switch r.Method {
	case http.MethodGet:
		if r.URL.Query().Has("id") {
			productId, e := strconv.Atoi(r.URL.Query().Get("id"))
			if e != nil {
				log.Println(e)
				response.ResponseError(w, http.StatusBadRequest, e)
				return
			}
			productApi.getProductById(w, r, productId)
			return
		}
		if r.URL.Query().Has("brand_id") {
			brandId, e := strconv.Atoi(r.URL.Query().Get("brand_id"))
			if e != nil {
				log.Println(e)
				response.ResponseError(w, http.StatusBadRequest, e)
				return
			}
			productApi.getProductsByBrand(w, r, brandId)
			return
		}
	case http.MethodPost:
		productApi.postProduct(w, r)
		return
	}
}
