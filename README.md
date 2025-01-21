# Design Patterns in Go

This repository is a comprehensive collection of Structural, Creational, and Behavioral Design Patterns, perfect for developers aiming to deepen their understanding of software design principles.

## Decorator in Go

Decorator is a ```structural pattern``` that allows adding new behaviors to objects dynamically by placing them inside special wrapper objects, called decorators.

## Conceptual Example
#### pizza.go: Component interface

```go
package main

type IPizza interface {
    getPrice() int
}
```
#### veggieMania.go: Concrete component
``` go
package main

type VeggieMania struct {
}

func (p *VeggieMania) getPrice() int {
    return 15
}
```
#### tomatoTopping.go: Concrete decorator
```go
package main

type TomatoTopping struct {
    pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
    pizzaPrice := c.pizza.getPrice()
    return pizzaPrice + 7
}
```
#### cheeseTopping.go: Concrete decorator
```go
package main

type CheeseTopping struct {
    pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
    pizzaPrice := c.pizza.getPrice()
    return pizzaPrice + 10
}
```
#### main.go: Client code
```go
package main

import "fmt"

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
}
```
#### output.txt: Execution result
``` 
Price of veggeMania with tomato and cheese topping is 32
```
