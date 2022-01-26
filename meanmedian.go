package main

import (
	"encoding/json"
	"fmt"
)

type meanmedian struct {
	Number float64 `json:"number"`
}

type o struct {
	Mean   float64 `json:"mean"`
	Median float64 `json:"median"`
}

// key left, velue right on num

var num = `[{"number":1},{"number":2},{"number":5},{"number":4}]`

func m() {
	var inp []meanmedian
	var out o
	json.Unmarshal([]byte(num), &inp)
	sum := 0.0
	fmt.Println(inp)
	// sort.Slice(inp, func(i, j int) bool {
	// 	return inp[i].Number < inp[j].Number
	// })
	// fmt.Println(inp, "ini sort")
	for i := 0; i < len(inp); i++ {
		sum = sum + inp[i].Number
	}
	mean := sum / float64(len(inp))
	i := 0
	median := 0.0
	if len(inp)%2 == 0 { // genap
		j := len(inp) / 2
		median = (inp[j].Number + inp[j-1].Number) / 2.0
	} else { // ganjil
		i = len(inp) / 2
		median = inp[i].Number
	}
	out.Mean = mean
	out.Median = median
	r, _ := json.Marshal(out)

	// fmt.Println(inp[sum])
	fmt.Println(sum)
	fmt.Println(mean)
	fmt.Println(median)
	fmt.Println(string(r))

}
