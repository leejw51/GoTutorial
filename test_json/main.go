package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := make(map[string]interface{})
	data["name"] = "apple"
	data["price"] = 1000
	data["apples"] = map[string]int{"apple": 5, "lettuce": 7}
	data["fruits"] = []string{"strawberry", "mango", "pear"}
	doc, _ := json.Marshal(data)
	fmt.Println(string(doc))

}
