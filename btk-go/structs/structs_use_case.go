package structs

import "fmt"

func StructUseCase() {
	firstProduct := product{
		name:         "Bilgisayar",
		unitPrice:    200,
		brand:        "Acer",
		discountRate: 5,
	}

	firstProduct.name = "Computer"
	firstProduct.unitPrice = firstProduct.IncreaseUnitPrice(200)
	fmt.Println(firstProduct.unitPrice)
}

type product struct {
	name         string
	unitPrice    float64
	brand        string
	discountRate int
}

func (p *product) IncreaseUnitPrice(amount float64) float64 {
	return p.unitPrice + amount
}
