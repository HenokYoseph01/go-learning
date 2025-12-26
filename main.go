package main

import (
	"fmt"
	"strings"
)

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
	scoreThree:=2.3 //defaults to float64

	fmt.Println(scoreOne, scoreTwo, scoreThree);

	//Arrays
	//1. Long way
	// var ages [3]int = [3]int{20, 30, 40}

	//2. Short way
	var ages = [3]int{50, 60, 70}

	fmt.Println(ages, len(ages)) //Use len() to get length of arrays

	//3. Short hand way (Shorter)
	names := [4]string{"Holy","Molly","This","Hard"}

	fmt.Println(names)

	// Slice (uses array under the hood), basically an array without a length specfication
	/* Can't append elements to an array, but can to a slice. 
	   append(<slice>, <newvalue>) returns a new slice, doesn't update the slice it references 
	   hence assign it to the reference slice
	*/

	var scores = []int{40,60,30};
	scores[2] = 50

	scores = append(scores, 80)

	fmt.Println(scores, len(scores))

	// Slice Ranges: Get range of elements from existing slice and assign it to a new slice
	rangeOne := names[1:3]
	rangeTwo := names[0:]
	rangeThree := names[:2]

	rangeOne = append(rangeOne, "That")

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	//December 26: Packages
	//strings package
	greetings := "Whats Up Gang"

	fmt.Println(strings.Contains(greetings,"Whats")) //Returns true or false if string exists in a string or not
	fmt.Println(strings.ReplaceAll(greetings, "Whats", "Hands")) //Replace a specfic string with another string (if string that is being searched to be replaced doesn't exist, return original string)
	fmt.Println(strings.ToUpper(greetings)) //Changes string to upper case
	fmt.Println(strings.Index(greetings, "ga"))
	fmt.Println(strings.Split(greetings, " ")) //Returns an array/splice with string split up based on a delimiter

	














}