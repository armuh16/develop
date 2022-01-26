package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	No               int    `json:"no"`
	LocalTransaction string `json:"localtransaction"`
	TotalAmount      int    `json:"total"`
}

type ResponResult struct {
	LocalTransaction string `json:"localtransaction"`
	TotalAmount      int    `json:"total"`
}

var inputJSON = `[
    {"no": 1, "localtransaction": "FIT SUNTER ALTERRA", "total": 3000},
    {"no": 2, "localtransaction": "FIT SUNTER ALTERRA", "total": 3000},
    {"no": 3, "localtransaction": "FIT PANCORAN", "total": 3000},
    {"no": 4, "localtransaction": "FIT BENHIL", "total": 3000},
    {"no": 5, "localtransaction": "FIT SUNTER ALTERRA", "total": 3000},
    {"no": 6, "localtransaction": "FIT ARTERI PONDOK INDAH", "total": 3000},
    {"no": 7, "localtransaction": "FIT BENHIL", "total": 3000},
    {"no": 8, "localtransaction": "FIT PANCORAN", "total": 3000},
    {"no": 9, "localtransaction": "FIT BENHIL", "total": 3000},
    {"no": 10, "localtransaction": "FIT BENHIL", "total": 3000}
  ]`

func main() {
	produk := []Product{}
	json.Unmarshal([]byte(inputJSON), &produk)
	listProduct := map[string]int{}
	for i := 0; i < len(produk); i++ {
		listProduct[produk[i].LocalTransaction] += produk[i].TotalAmount
	}
	result := make([]ResponResult, len(listProduct))
	var i int
	for key, value := range listProduct {
		result[i].LocalTransaction = key
		result[i].TotalAmount = value
		i++
	}
	jsonData, _ := json.Marshal(result)
	// fmt.Println(string(inputJSON))
	fmt.Println(string(jsonData))
}
