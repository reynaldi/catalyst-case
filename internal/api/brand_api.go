package api

import (
	"catalyst-case/database"
	"catalyst-case/module/brand"
	"catalyst-case/pkg/response"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type brandApi struct {
	brand brand.Brand
}

func newBrandApi(db *database.DB) *brandApi {
	brandQuery := brand.NewBrandQuery(db)
	brandCommand := brand.NewBrandCommand(db)
	return &brandApi{
		brand: brand.NewBrand(brandQuery, brandCommand),
	}
}

func (b *brandApi) getAllBrands(w http.ResponseWriter, r *http.Request) {
	res, err := b.brand.GetAllBrands(r.Context())
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusNoContent, err)
		return
	}
	response.ResponseWithJSON(w, http.StatusOK, res)
}

func (b *brandApi) getBrandById(w http.ResponseWriter, r *http.Request, brandId int) {
	res, err := b.brand.GetBrandById(r.Context(), brandId)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusNoContent, err)
		return
	}
	response.ResponseWithJSON(w, http.StatusOK, res)
}

func (b *brandApi) postBrand(w http.ResponseWriter, r *http.Request) {
	var model *brand.NewBrandDto
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusNoContent, err)
		return
	}
	validate := validator.New()
	err = validate.Struct(model)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusNoContent, err)
		return
	}
	err = b.brand.AddNewBrand(r.Context(), *model)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusNoContent, err)
		return
	}
	response.ResponseWithNoContent(w)
}
