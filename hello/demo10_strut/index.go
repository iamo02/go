package main

import "fmt"

type product struct {
	name   string
	prince int
	stock  int
}

func main() {
	var p1 product
	p1.name = "vue"
	p1.prince = 100
	p1.stock = 1

	show(p1)
	update(&p1)
	show(p1)

	p1 = p1.clear()

	show(p1)
}

func (p product) clear() product {
	p.prince = 0
	p.stock = 0
	return p
}

func show(p product) {
	fmt.Println(p.prince)
}

func update(p *product) {
	p.prince = p.prince + 200
}
