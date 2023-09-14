package main

import (
	"fmt"
	"reflect"
)

func main() {
	variablesTypes()
}

func variablesTypes() {
	var name string = "Junior Raveline"
	var age int8 = 27
	var balance float32 = 6280.89
	isMarried := false
	goodLookingScale := 9.7

	fmt.Println("My name is", name)
	fmt.Println("My balance on bank is $", balance, "and my age is", age, "years old.")

	fmt.Println(name, "is married?", isMarried)
	fmt.Println(name, "good looking scale is", goodLookingScale)
	fmt.Println("The variable goodLookingScale is the type is", reflect.TypeOf(goodLookingScale))
}
