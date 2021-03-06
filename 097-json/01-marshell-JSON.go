// Marshell to JSON is a procedure of converting a Go Struct into a JSON
package main

import (
	"fmt"
	"encoding/json"
)

type employee struct {
	FirstName string // should be capital in order to work
	LastName string
}

func main() {
	fmt.Println("Marshell JSON")

	abhi := employee{
		FirstName: "abhishek",
    LastName: "das",
	}

	su := employee{
		FirstName: "sutapa",
		LastName: "chattaraj",
	}

	malpa := employee{
		FirstName: "malpa",
		LastName: "bajpai",
	}

	xe := []employee{abhi,su,malpa}
	
	fmt.Println("Go data:", xe)// OUTPUT: [{abhishek das} {sutapa chattaraj} {malpa bajpai}]
	fmt.Printf("Go formatted data: %+v\n", xe) // OUTPUT: [{abhishek das} {sutapa chattaraj} {malpa bajpai}]
	// convert into JSON
	bs, err:= json.Marshal(xe) //bs stands for byte of strings
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("JSON data:", string(bs)) //{"FirstName":"abhishek","LastName":"das"},{"FirstName":"sutapa","LastName":"chattaraj"},{"FirstName":"malpa","LastName":"bajpai"}]
}