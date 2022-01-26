package main

import (
	"encoding/json"
	"fmt"
)

type n struct {
	Number1 int `json:"n1"`
	Number2 int `json:"n2"`
	Sum     int `json:"result"`
}

var result = `{"n1":4, "n2":2}`

func main() {
	var number n
	json.Unmarshal([]byte(result), &number)
	sum := number.Number1 / number.Number2
	number.Sum = sum

	r, _ := json.Marshal(number)
	fmt.Println(string(r))
	// fmt.Println(number)

}
