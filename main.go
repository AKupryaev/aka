package main

import (
	"errors"
	"fmt"
)

// Структура товара
type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

// Структура склада
type Inventory struct {
	products map[int]Product
}

// Конструктор
func NewInventory() *Inventory {
	return &Inventory{
		products: make(map[int]Product),
	}
}

// Добавление товара
func (inv *Inventory) AddProduct(p Product) {
	inv.products[p.ID] = p
}

// Продажа товара
func (inv *Inventory) SellProduct(id, qty int) error {
	product, ok := inv.products[id]
	if !ok {
		return errors.New("товар не найден")
	}

	if product.Quantity < qty {
		return errors.New("недостаточно товара")
	}

	product.Quantity -= qty
	inv.products[id] = product // важно обновить!

	return nil
}

func main() {
	inv := NewInventory()

	inv.AddProduct(Product{ID: 1, Name: "Телефон", Price: 30000, Quantity: 10})

	fmt.Println(inv.SellProduct(1, 3)) // nil
	fmt.Println(inv.SellProduct(1, 8)) // ошибка

	fmt.Println(inv.products[1]) // проверка остатка
}
