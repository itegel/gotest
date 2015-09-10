package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type ColorGroup struct {
		ID     int    `json:"id"`
		Name   string `json:"cookiename"`
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%s\n", b)

	var group2 ColorGroup

	json.Unmarshal(b, &group2)

	fmt.Printf("%#v\n", group2)
}
