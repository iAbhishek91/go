/* demonstration of basic value based interface */
package main

import "fmt"

// 1. define a interface
type cart interface {
	print()
	addItem()
	deleteItem()
}

// 2. define a struct
type amazon struct {
	userName string
	items uint8
}

// 3. implement all the function of the interface
func (a amazon) print() {
	fmt.Printf("Hi %s, you have %d items in your cart!\n", a.userName, a.items)
}
func (a *amazon) addItem() {
	a.items ++
	fmt.Println("New items added to your cart!")
}
func (a *amazon) deleteItem() {
	a.items --
	fmt.Println("Deleted an item from your cart!")
}



func main() {
	// method sets
	// we have both value and pointer receiver
	// hence as per method set rules we need pass point variable.
	myCart := &amazon{"abhishek", 3}
	printCart(myCart)

	addItemToCart(myCart)
	printCart(myCart)

	deleteItemToCart(myCart)
	printCart(myCart)
}


func printCart(c cart) {
	c.print()
}

func addItemToCart(c cart) {
	c.addItem()
}

func deleteItemToCart(c cart) {
	c.deleteItem()
}