package main

import (
	"fmt"
	"os"
)

// bill structure
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newbill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// bill format
func (b bill) format() string {
	fs := "bill breakdown: \n"
	var total float64 = 0
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v....%v\n", k+":", +v)
		total += v
	}
	fs += fmt.Sprintf("%-25v....%0.2f\n", "tip:", b.tip)
	fs += fmt.Sprintf("%-25v....%0.2f", "total:", total+b.tip)
	return fs
}

// update tip
func (b *bill) updatetip(tip float64) {
	b.tip = tip
}

//update  items
func (b bill) additems(name string, price float64) {
	b.items[name] = price

}

// save bill
func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill saved to file")
}
