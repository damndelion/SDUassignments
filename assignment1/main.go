package main

import (
	"fmt"
	entity2 "github.com/damndelion/SDUassignments/assignment1/entity"
	"github.com/damndelion/SDUassignments/assignment1/utils"
)

func main() {
	var x int = 10          // x := 10
	var y string = "Daniar" // y := "Daniar"

	fmt.Println(x)
	fmt.Println(y)

	utils.EvenOrOdd(10)
	utils.OneToFive()

	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[1:3]
	fmt.Println(s)

	personInfo := make(map[string]string)
	personInfo["name"] = "Daniar"
	personInfo["age"] = "20"
	personInfo["city"] = "Almaty"

	fmt.Println("Name", personInfo["name"])
	fmt.Println("Age", personInfo["age"])
	fmt.Println("City", personInfo["city"])

	p1 := entity2.Person{Name: "Daniar", Age: 20, City: "Almaty"}
	p2 := entity2.Person{Name: "Joe", Age: 50, City: "London"}

	p1.PrintInfo()
	p2.PrintInfo()

	e1 := entity2.Employee{Person: entity2.Person{Name: "Daniar", Age: 20, City: "Almaty"}, JobTitle: "Unemployed("}

	entity2.DisplayInfo(p1)
	entity2.DisplayInfo(e1)

}
