package entity

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) PrintInfo() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("City:", p.City)
}

func (p Person) Print() {
	p.PrintInfo()
}
