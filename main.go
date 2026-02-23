package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

// Feb 10 Package Scope
var score = 99.5

// Feb 5 Functions
func sayGreeting(n string){
	fmt.Printf("Good Morning %v \n", n)
}

func sayBye(n string){
	fmt.Printf("Good Bye %v \n", n)
}

func cycleNames(n []string, f func(string)){
	for _, value := range n{
		f(value);
	}
}

func circleArea(r float64) float64{
	return math.Pi * r * r
}

// Feb 9 Multiple returns
func getInital(n string) (string, string){
	s := strings.ToUpper(n);
	names := strings.Split(s, " ");

	var initals []string

	for _, v := range names {
		initals = append(initals, v[:1])
	}

	if len(initals) > 1 {
		return initals[0], initals[1]
	}

	return initals[0], "_"

}
//Feb 17: Pass By Value
func updateName(n string) string{
		n = "Change"
		return n;
}

func updateMenu(y map[string]float64) {
	y["coffee"]=2.99
}


//Feb 18: Pointers
//Accept pointer as parameter via *<type>
func updateNamePointer(n *string){
	*n = "Change Via Pointer"
}

//Feb 20: User input

//Helper function
func getInput (prompt string, r *bufio.Reader) (string, error){
	fmt.Print(prompt)
	 input, err := r.ReadString('\n') //Read after new line (after enter has been clicked)
	 return strings.TrimSpace(input), err

}

func createBill() bill {
	 reader := bufio.NewReader(os.Stdin)

	//  fmt.Print("Create a new bill name: ")
	//  name, _ := reader.ReadString('\n') //Read after new line (after enter has been clicked)
	//  name = strings.TrimSpace(name);

	name, _ := getInput("Create a new bill name: ", reader)
	 b := newBill(name);
	 fmt.Println("Create a Bill - ", b.name);

	 return b;
}

func promptOptions(b bill){
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose Option (a - add item, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)

 
}

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

	//Original String doesn't get altered
	fmt.Println("Original String:", greetings) 

	//sort package
	numbers := []int{45, 20, 42, 90, 39, 10, 20, 28}

	sort.Ints(numbers) //sorts package alters original slice
	fmt.Println(numbers)

	int_index := sort.SearchInts(numbers, 90)
	fmt.Println(int_index) //Returns index of number that wants to be searched, if not there then will return element + 1 (it doesn't exist)

	name_slice := []string{"Some","Strings","To","Be","Arranged","Type"}
	sort.Strings(name_slice)
	fmt.Println(name_slice)
	string_index := sort.SearchStrings(name_slice, "To")
	fmt.Println(string_index)


	//February 5: Loops
	x := 0;
	for x < 5 {
		fmt.Println("Value of x is:", x)
  		x++;
	}

	for i:= 0; i<5; i++ {
		fmt.Println("Value of i is:", i)

	}

	nameLoop := []string{"mario","luigi","yoshi","peach"}
	for i:=0; i<len(nameLoop); i++{
		fmt.Println(nameLoop[i])
	}

	//Modern loop by using range
	for index, value := range nameLoop {
		fmt.Printf("The value at index %v is %v \n", index, value)
		value = "new stuff" //Doesn't change value of array at corresponding index as value is like a local variable
	}

	fmt.Println(nameLoop);

	//Lesson: Boolean
	age := 45
	fmt.Println(age<=50)
	fmt.Println(age>=50)
	fmt.Println(age==50)
	fmt.Println(age!=50)

	if age < 30 {
		fmt.Println("age is less than 30")
	}else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("Age is not less than 45")
	}

	for index, value := range nameLoop {
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("Breaking at pos", index)
			break
		}

		fmt.Printf("The value at index %v is %v \n", index, value)
	}

	//FUNCTIONS
	// sayGreeting("Dozy");
	// sayGreeting("Viby");
	// sayBye("Moonz");
	cycleNames([]string{"Smiley","Psycho","Viby"}, sayGreeting)
	a1 := circleArea(10.2)
	a2 := circleArea(18)

	fmt.Println(a1, a2);
	fmt.Printf("Circle 1 is %0.3f and Circle 2 is %0.3f", a1, a2)
	fmt.Println("")

	//Feb 9: Multiple Returns
	fn1, sn1 := getInital("John Snow")
	fmt.Println(fn1, sn1)
	fn2, sn2 := getInital("Scarlet Rot")
	fmt.Println(fn2, sn2)
	fn3, sn3 := getInital("Melinia")
	fmt.Println(fn3, sn3)

	//Feb 10: Package Scope (greetings.go)
	//To make sayHello and the for range loop on the points slice work, run both main and greeting files (go run main.go greetings.go)
	// sayHello("mario")

	// for _ , value := range points{
	// 	fmt.Println(value)
	// }

	// showScore();

	//Feb 10: Maps
	//Keys and values have to be of the same type (ex: integer keys and string values)
	//syntax: map[<keyType>]<valueType>{key:value}
	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"]) //Value based on key/ specfic value from map

	//Looping maps
	for key, value := range menu {
		fmt.Println(key, "-",value)
	}

	// ints as key type
	phoneBook := map[int]string{
		11111: "Kirby",
		22222: "Link",
		33333: "Sonic",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[11111])

	//Update value
	phoneBook[22222] = "Zelda"
		fmt.Println(phoneBook)

	//Feb 17: Pass by Value
	
	//Group A
	name:="Hi"
	name=updateName(name)
	fmt.Println(name)

	//Group B
	menus:= map[string]float64{
		"pie": 4.33,
		"ice cream": 2.33,
	}

	updateMenu(menus)

	fmt.Println(menus)

	//Feb 18: Pointers

	n:= "Heyo" 
	//Pointer to memory location
	fmt.Println("Memory Address of variable name: ", &n)

	m:= &n;
	fmt.Println("memory address: ", m)
	//Dereference with *
	fmt.Println("Value at memory address: ", *m)
	fmt.Println("Before update: ", n)
	updateNamePointer(m)
	fmt.Println("Updated name: ", n)


	//Feb 19: Struct & Custom Types
	myBill := newBill("Broke Boy's Bill")
	fmt.Println(myBill);

	//Reciever function
	fmt.Println(myBill.format())

	//Reciever function (pointer)
	myBill.addItem("Pizza", 5.99)
	myBill.addItem("Coke", 1.99)
	myBill.updateTip(10)
	fmt.Println(myBill.format())

	//Feb 20: User Input
	mybill2 := createBill()
	promptOptions(mybill2)
	fmt.Println(mybill2)

	

}
