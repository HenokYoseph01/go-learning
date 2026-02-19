package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Function to make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 3.99, "doughnut": 2.29},
		tip:   0,
	}

	return b

}

// Format bill (reciever function)
// Set a receiver by stating in a paranthesis after the func keyword the struct it can only run on
// When variable with struct calls this function, b only takes the copy of the struct
func (b bill) format() string {
	fs := "Bill breakdown \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f","total:", total);

	return fs;
}