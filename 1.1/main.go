package main

import "fmt"

type Human struct {
	health int
	Name string
}

func (h *Human) Damage(hp int) {
	h.health -= hp
}

func (h Human) Describe() {
	fmt.Printf("This is %s with %d HP\n", h.Name, h.health)
}

type Action struct {
	Human
	Type string
}

func (a *Action) Kill() {
	a.health = 0
}

func main() {
	a := Action{
		Human: Human{
			health: 100,
			Name: "John",
		},
		Type: "Container",
	}	

	fmt.Println("I can access fields of embedded struct directly. Name:", a.Name)
	fmt.Println("I can access fields of embedded struct with full path. Name:", a.Human.Name)
	
	// Call methods of embedded struct directly
	a.Damage(10)
	a.Describe()
	// Call methods of embedded struct with full path
	a.Human.Damage(10)
	a.Human.Describe()
	
	fmt.Println("Still can access container's fields:", a.Type)
	a.Kill()

	type Humanable interface {
		Describe()
	}

	// Can use container type as implementation of embedded struct interface
	var h Humanable = a
	h.Describe()
}