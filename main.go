package main

import "fmt"

func main() {
	fmt.Println("Hello, Peeps")

	//Strings 
	var nameOne string = "Midoria"
	var nameTwo = "Bakugo"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameThree = "Todoroki"

	fmt.Println(nameOne, nameTwo, nameThree)

	//Short variable declaration, once declared, no need to := when changing value. Can only be used in functions, and not outside function scope
	nameFour := "Uraraka"
	fmt.Println(nameFour)

	//Ints
	var ageOne int = 24
	var ageTwo = 25
	ageThree := 26

	fmt.Println(ageOne, ageTwo, ageThree)

	//Int with bits
	var specAge int8 = 127
	var specAgeTwo int16 = -32000
	var specAgeThree uint8 = 255

	fmt.Println(specAge, specAgeTwo, specAgeThree)

	//Floats
	var scoreOne float32 = 3.4
	var scoreTwo float64 = 2000.21
	scoreThree:=2.3 //defaults to float65

	fmt.Println(scoreOne, scoreTwo, scoreThree);





}