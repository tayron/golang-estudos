package main

import "fmt"

// Strategy define a interface para as estratégias de cálculo de preço
type Strategy interface {
	CalculatePrice(float64) float64
}

// NormalStrategy implementa a estratégia de cálculo de preço normal
type NormalStrategy struct{}

func (s *NormalStrategy) CalculatePrice(price float64) float64 {
	return price
}

// DiscountStrategy implementa a estratégia de cálculo de preço com desconto
type DiscountStrategy struct {
	discountPercentage float64
}

func (s *DiscountStrategy) CalculatePrice(price float64) float64 {
	return price * (1 - s.discountPercentage/100)
}

// ShoppingCart é a estrutura que possui uma estratégia de cálculo de preço e itens
type ShoppingCart struct {
	strategy Strategy
	items    []float64
}

func (cart *ShoppingCart) SetStrategy(strategy Strategy) {
	cart.strategy = strategy
}

func (cart *ShoppingCart) AddItem(price float64) {
	cart.items = append(cart.items, price)
}

func (cart *ShoppingCart) CalculateTotal() float64 {
	total := 0.0
	for _, price := range cart.items {
		total += cart.strategy.CalculatePrice(price)
	}
	return total
}

func main() {
	normalStrategy := &NormalStrategy{}
	discountStrategy := &DiscountStrategy{discountPercentage: 10}

	cart := ShoppingCart{}
	cart.AddItem(100)
	cart.AddItem(50)
	cart.AddItem(75)

	// Use a estratégia de cálculo de preço normal
	cart.SetStrategy(normalStrategy)
	fmt.Printf("Total com preço normal: %.2f\n", cart.CalculateTotal())

	// Use a estratégia de cálculo de preço com desconto
	cart.SetStrategy(discountStrategy)
	fmt.Printf("Total com desconto: %.2f\n", cart.CalculateTotal())
}
