package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	name string
	age  int
}

func (h *Human) GrownUp() {
	h.age++
}

type Action struct {
	Human
}

func (a Action) SayName() {
	fmt.Printf("My name is %s\n", a.Human.name)
}

func main() {
	//создаем объект Action
	person := Action{Human{
		name: "Petya",
		age:  10,
	}}

	//используем метод Action
	person.SayName()

	//используем метод Human
	person.GrownUp()

	fmt.Println(person.Human.age)
}
