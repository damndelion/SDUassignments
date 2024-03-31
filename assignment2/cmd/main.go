package main

import (
	"fmt"
)

func main() {
	//app.Run()

	x := [2]string{"3", "4"}
	for y := range x {
		fmt.Println(y)
	}
	var val interface{} = "foo"
	if str, ok := val.(string); ok {
		fmt.Println(str)
	}
}
