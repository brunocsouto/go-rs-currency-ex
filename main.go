package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type FileData struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 2 {
		panic("Please provide the value in BRL and the destination currency as arguments.")
	}
	moeda_destino := argsWithoutProg[1]
	valor_em_brl := argsWithoutProg[0]
	floatValue, err := strconv.ParseFloat(valor_em_brl, 32)
	check(err)

	data, err := os.ReadFile("./rates.json")
	check(err)
	byt := []byte(data)

	var fileData FileData
	err = json.Unmarshal(byt, &fileData)
	check(err)

	rate := fileData.Rates[moeda_destino]
	result := floatValue * rate
	fmt.Printf("%.2f\n", result)
}
