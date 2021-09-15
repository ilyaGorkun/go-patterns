package adapter

import "testing"

func TestAdapter(t *testing.T) {
	pizza:= &veggeMania{}

	if pizza.getPrice() != 15 {
		t.Error("Expected 15, got", pizza.getPrice())
	}

	pizzaWithTomato := &tomatoTopping{pizza: pizza}

	if pizzaWithTomato.getPrice() != 22 {
		t.Error("Expected 22, got", pizzaWithTomato.getPrice())
	}

	pizzaWithCheese := &cheeseTopping{pizza: pizza}

	if pizzaWithCheese.getPrice() != 25 {
		t.Error("Expected 25, got", pizzaWithCheese.getPrice())
	}

	pizzaWithTomatoAndCheese := &tomatoTopping{
		pizza: &cheeseTopping{pizza: pizza},
	}


	if pizzaWithTomatoAndCheese.getPrice() != 32 {
		t.Error("Expected 32, got", pizzaWithTomatoAndCheese.getPrice())
	}
}
