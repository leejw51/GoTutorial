package main

import (
	"encoding/json"
	"fmt"
)

func test1() {
	data := make(map[string]interface{})
	data["name"] = "apple"
	data["price"] = 1000
	data["apples"] = map[string]int{"apple": 5, "lettuce": 7}
	data["fruits"] = []string{"strawberry", "mango", "pear"}
	data["box"] = make([]int, 10, 100)
	data["numbers"] = []int{10, 20, 30}
	doc, _ := json.Marshal(data)
	fmt.Println(string(doc))
}

func test2() {
	data := make([]interface{}, 0)
	data = append(data, 4)
	data = append(data, "0x1c")
	data = append(data, []int{25, 75})
	doc, _ := json.Marshal(data)
	fmt.Println(string(doc))

}

func main() {
	data := make([]interface{}, 0)
	data = append(data, 4)
	data = append(data, "0x1c")
	data = append(data, []int{25, 75})
	doc, _ := json.Marshal(data)

	var d struct {
		a int    `json:"test"`
		b string `json:"TEST"`
		c []int  `json`
	}
	_ := json.Unmarshal(doc, &d)
	data2, _ := json.Marshal([]int{25, 75})
	fmt.Println(string(doc))
	fmt.Println(string(data2))
}
