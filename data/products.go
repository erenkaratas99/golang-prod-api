package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"regexp"
	"time"
)

// products defined for the product API
// swagger:model
type Product struct {
	// the id of the user
	//
	// required : true
	// min :1
	ID        int    `json:"id"`
	Name      string `json:"name" validate:"required"`
	Amount    int    `json:"amount" validate:"gt=0"`
	SKU       string `json:"sku" validate:"required,sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
}

type Products []*Product //custom type

func GetProducts() Products {
	return productList
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int { //mock db working principle implementation
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")
var ErrMoreThanOneProd = fmt.Errorf("More than one product with same ID")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func findProduct4DeleteReq(id int) (*Product, int, error) {
	var feedbackList []int
	for i, p := range productList {
		if p.ID == id {
			feedbackList = append(feedbackList, i)
		}
	}
	if len(feedbackList) == 0 {
		return nil, -1, ErrProductNotFound
	}
	if len(feedbackList) == 1 {
		return nil, feedbackList[0], nil
	}
	return nil, -1, ErrMoreThanOneProd

}

func DeleteProduct(id int) error {
	_, i, _ := findProduct4DeleteReq(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1])

	return nil
}

var productList = []*Product{
	&Product{
		ID:        1,
		Name:      "product-1",
		Amount:    10,
		SKU:       "abc-dce-fgh",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID:        2,
		Name:      "product-2",
		Amount:    20,
		SKU:       "eren-tesodev-itu",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
