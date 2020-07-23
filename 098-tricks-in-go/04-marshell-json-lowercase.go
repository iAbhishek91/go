// Marshell to JSON is a procedure of converting a Go Struct into a JSON
package main

import (
	"fmt"
	"encoding/json"
)

type employee struct {
	firstName string
	lastName string
}

func main() {
	fmt.Println("Marshell JSON")

	abhi := employee{
		firstName: "abhishek",
    lastName: "das",
	}

	su := employee{
		firstName: "sutapa",
		lastName: "chattaraj",
	}

	malpa := employee{
		firstName: "malpa",
		lastName: "bajpai",
	}

	xe := []employee{abhi,su,malpa}
	
	fmt.Println("Go data:", xe)// OUTPUT: [{abhishek das} {sutapa chattaraj} {malpa bajpai}]
	fmt.Printf("Go formatted data: %+v\n", xe) // OUTPUT: [{abhishek das} {sutapa chattaraj} {malpa bajpai}]
	// convert into JSON
	bs, err:= json.Marshal(xe) //bs stands for byte of strings
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("JSON data:", string(bs)) // empty as fields need to be exported
}