package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getinput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err

}
func createbill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getinput("creating new bill:", reader)
	b := newbill(name)
	fmt.Println("created new bill:", b.name)
	return b

}
func promptopt(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getinput("choose options (a.....additem,b.....add tip,c......save bill:", reader)
	switch opt {
	case "a":
		name, _ := getinput("item name:", reader)
		price, _ := getinput("item price:", reader)
		f, err := strconv.ParseFloat(price, 64)
		b.items[name] = f
		if err != nil {
			fmt.Println("the price must be in number")
			promptopt(b)

		}
		fmt.Println(name, price)
		promptopt(b)
	case "b":

		tip, _ := getinput("enter the tip amount($):", reader)
		p, err := strconv.ParseFloat(tip, 64)
		b.tip = p
		if err != nil {
			fmt.Println("the tip must be in number")
			promptopt(b)
		}

		fmt.Println(tip)
		promptopt(b)
	case "c":
		b.save()
		fmt.Println("you chose to save the bill", b)
	default:
		fmt.Println("invalid option")
		promptopt(b)
	}

}

func main() {
	mybill := createbill()
	promptopt(mybill)

}
