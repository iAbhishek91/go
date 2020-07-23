package main

import (
	"fmt"
	"encoding/json"
)

type employee struct {
	FirstName string
	LastName string
}

func main() {
	j := `[{"FirstName":"abhishek","LastName":"das"},{"FirstName":"sutapa","LastName":"chattaraj"},{"FirstName":"malpa","LastName":"bajpai"}]`

	s := []employee{}

	err := json.Unmarshal([]byte(j), &s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", s)
}