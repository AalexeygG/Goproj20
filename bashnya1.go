package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type myStruct []struct {
	GlobalId int64 `json:"global_id"`
}

func main1() {
	file, err := os.Open("data-20190514T0100.json")
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var ans myStruct
	if err := json.Unmarshal(data, &ans); err != nil {
		panic(err)
	}
	var sum int64
	for i := range ans {
		sum += ans[i].GlobalId
	}
	fmt.Println(sum)
}
