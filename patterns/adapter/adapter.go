package adapter

type pizza interface {
	getPrice() int
}

type veggeMania struct {
}

func (p *veggeMania) getPrice() int {
	return 15
}

type tomatoTopping struct {
	pizza pizza
}

func (t *tomatoTopping) getPrice() int {
	return 7 + t.pizza.getPrice()
}

type cheeseTopping struct {
	pizza pizza
}

func (t *cheeseTopping) getPrice() int {
	return 10 + t.pizza.getPrice()
}
