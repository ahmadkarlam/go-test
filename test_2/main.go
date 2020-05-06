package main

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Product struct {
	name     string
	quantity int
}

type ProductCollection struct {
	products map[string]Product
	mux      sync.Mutex
}

func NewProductCollection() ProductCollection {
	return ProductCollection{
		products: map[string]Product{
			"Baju": {
				name:     "Baju",
				quantity: 1,
			},
			"Kemeja": {
				name:     "Kemeja",
				quantity: 5,
			},
		},
	}
}

func (p *ProductCollection) Buy(name string) error {
	p.mux.Lock()
	product := p.products[name]
	if product.quantity == 0 {
		p.mux.Unlock()
		return errors.New("out of stock")
	}
	product.quantity -= 1
	p.products[name] = product
	p.mux.Unlock()
	return nil
}

var Products = NewProductCollection()

func main() {
	r := gin.Default()
	r.GET("/order", func(ctx *gin.Context) {
		productName := ctx.Query("product_name")
		if err := orderProduct(Products, productName); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("something went wrong: %s", err),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("success buy %s", productName),
		})
	})

	if err := r.Run(); err != nil {
		panic(err)
	}
}

func orderProduct(collection ProductCollection, name string) error {
	if _, ok := collection.products[name]; !ok {
		return errors.New("product not found")
	}
	if err := Products.Buy(name); err != nil {
		return err
	}

	return nil
}
