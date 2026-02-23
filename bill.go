package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Function to make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b

}

// Format bill (reciever function)
// Set a receiver by stating in a paranthesis after the func keyword the struct it can only run on
// When variable with struct calls this function, b only takes the copy of the struct
func (b *bill) format() string {
	fs := "Bill breakdown \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	//Add Tio
	fs += fmt.Sprintf("%-25v ...$%v\n","tip:", b.tip)

	//Total
	fs += fmt.Sprintf("%-25v ...$%0.2f","total:", total+b.tip);

	return fs;
}

// Update the tip (via pointer)
//Reciever function having pointers is advantegous as it saves up space once created rather than creating it normally and having a copy made everytime in memory when called
func (b *bill) updateTip(tip float64) {
	b.tip = tip  //automatically dereferenced because structs are pointer types
}

// Add item to bill
func (b *bill) addItem(name string, price float64){
	b.items[name] = price;
}

//Feb 23: Files
func (b *bill) save(){
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil{
		panic(err)
	}

	fmt.Printf("Bill was saved to file")
}