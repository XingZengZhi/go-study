package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Address   []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	//将数据格式化（编码）为JSON格式（[]byte）
	js, _ := json.Marshal(vc)

	fmt.Printf("JSON format: %s", js)

	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)

	defer file.Close()

	enc := json.NewEncoder(file)

	//数据解码
	err := enc.Encode(vc)

	if err != nil {
		log.Println("Error in encoding json")
	}
}
