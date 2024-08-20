package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

type AnimalAdapter interface {
	Say()
}

type Fish struct {
}

type Cat struct {
}

func (f *Fish) Gurgling() {
	fmt.Println("Бульк...")
}

func (c *Cat) Meow() {
	fmt.Println("Мяу!")
}

type FishAdapter struct {
	*Fish
}

type CatAdapter struct {
	*Cat
}

func (adapter *FishAdapter) Say() {
	adapter.Gurgling()
}

func (adapter *CatAdapter) Say() {
	adapter.Meow()
}

func NewFishAdapter(fish *Fish) AnimalAdapter {
	return &FishAdapter{fish}
}

func NewCatAdapter(cat *Cat) AnimalAdapter {
	return &CatAdapter{cat}
}

func main() {
	zoo := [...]AnimalAdapter{NewCatAdapter(&Cat{}), NewFishAdapter(&Fish{})}

	for _, animal := range zoo {
		animal.Say()
	}

}
