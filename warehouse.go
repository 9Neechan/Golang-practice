package main

import "fmt"

type Product struct {
	ID       int
	Name     string
	Quantity int
}

type Storage struct {
	Products map[int]Product
}

type Warehouse interface {
	AddProduct(product Product) error
	UpdateQuantity(productID int, quantity int) error
}

func (st *Storage) AddProduct(product Product) error {
	if st.Products == nil {
		st.Products = make(map[int]Product)
	}

	if _, ok := st.Products[product.ID]; ok {
		return fmt.Errorf("товар с ID %d уже существует на складе", product.ID)
	}

	st.Products[product.ID] = product
	return nil
}

func (st *Storage) UpdateQuantity(productID int, quantity int) error {
	product, ok := st.Products[productID]
	if !ok {
		return fmt.Errorf("товар с ID %d не найден на складе", productID)
	}

	if product.Quantity + quantity < 0 {
		return fmt.Errorf("количество товара не может быть отрицательным")
	}

	product.Quantity += quantity
	st.Products[productID] = product
	return nil
}