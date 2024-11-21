package main

import "fmt"

// pizza.go
type IPizza interface {
	getPrice() int
}

// veggieMania.go
type VeggieMania struct {
}

func (p *VeggieMania) getPrice() int {
	return 15
}

// tomatoTopping.go
type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

// cheeseTopping.go
type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

// main.go
func main() {

	pizza := &VeggieMania{}

	//Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
	fmt.Println()
	fmt.Println()

	compMac := &Mac{}
	sumkalyComp := &Sumka{product: compMac}
	ativiruslyComp := &AntiVirus{product: sumkalyComp}
	fmt.Println(ativiruslyComp.getPrice())

	compHp := &Hp{}
	ativiruslyComp = &AntiVirus{product: compHp}
	fmt.Println(ativiruslyComp.getPrice())

}

type Product interface {
	getPrice() int
}

type Mac struct{}

func (m *Mac) getPrice() int {
	return 10
}

type Hp struct{}

func (m *Hp) getPrice() int {
	return 7
}

type Sumka struct {
	product Product
}

func (m *Sumka) getPrice() int {
	return m.product.getPrice() + 20
}

type AntiVirus struct {
	product Product
}

func (av *AntiVirus) getPrice() int {
	return av.product.getPrice() + 12
}
