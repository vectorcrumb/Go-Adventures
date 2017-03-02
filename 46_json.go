package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Page	int
	Fruits	[]string
}

type Response2 struct {
	Page	int	 `json:"page"`
	Fruits	[]string `json:"fruits"`
}

func main() {
	// To encode basic data types, we do the following
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))
	// The same applies for slices and maps
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	// The JSON package can automatically convert custom data types
	res1D := &Response1{
		Page:	1,
		Fruits:	[]string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	// You can also use tags on struct field declarations for custom JSON fields
	res2D := &Response2{
		Page:	1,
		Fruits:	[]string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
	// To decode JSON, we create a sample string for the example and later
	// a map to store the information. We also then check2 for errors.
	byt := []byte(`{"num":6.13, "strs":["a", "b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	// To use the values in a decoded map, we must cast them to their appropriate type
	num := dat["num"].(float64)
	fmt.Println(num)
	// We access nested data with multiple casts
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)
	// We can also decode JSON into custom data types. This provides additional type safety
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
	// We can also encode JSON directly into different streams
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}