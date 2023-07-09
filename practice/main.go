package main

import (
	"fmt"
)

// special type of variable constsnt
const aConst = 64 // can also be inferred like this, this is available to any function in the file

func main() {
	var aString string = "This is Go"
	fmt.Println(aString)
	fmt.Printf("The variable name is %T \n", aString) //%T is a verb or a placeholder
	// Beware Printf does not add a line by default you need to add it manually
	var anInteger int = 90
	fmt.Println(anInteger)

	var defaultInt int //all types have a defualt value for int it is 0
	fmt.Println(defaultInt)

	var anotherString = "This is a value of another string" // this isn't explicitly defined
	fmt.Printf("The variable name is %T \n", anotherString)

	// another style that you can use to initialise the varibale := no need to use var
	// this only works insdie functions, outside you must use the var keyword
	myString := "This is my string"
	fmt.Println(myString)

}
