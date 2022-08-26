package api

import (
	"catalyst-case/database"
	"catalyst-case/module/product"
	"catalyst-case/pkg/response"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type productApi struct {
	product product.Product
}

func newProductApi(db *database.DB) *productApi {
	productQuery := product.NewProductQuery(db)
	productCommand := product.NewProductCommand(db)
	return &productApi{
		product: product.NewProduct(productQuery, productCommand),
	}
}

func (p *productApi) postProduct(w http.ResponseWriter, r *http.Request) {
	var model *product.NewProductDto
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
	err = p.product.AddNewProduct(r.Context(), *model)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	response.ResponseWithNoContent(w)
}

func (p *productApi) getProductById(w http.ResponseWriter, r *http.Request, productId int) {
	res, err := p.product.GetProductById(r.Context(), productId)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	response.ResponseWithJSON(w, http.StatusOK, res)
}

func (p *productApi) getProductsByBrand(w http.ResponseWriter, r *http.Request, brandId int) {
	res, err := p.product.GetProductsByBrand(r.Context(), brandId)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	response.ResponseWithJSON(w, http.StatusOK, res)
}
