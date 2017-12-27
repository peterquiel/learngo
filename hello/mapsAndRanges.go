package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2

	fmt.Println("Map: ", m, " Length:", len(m))

	delete(m, "one")
	fmt.Println("Map: ", m, " Length:", len(m))

	// the blank identifier '_' ignores a return value
	_, presentInMap := m["one"] // second optional return value indicates if the key is present in the,

	fmt.Println(presentInMap)

	n := map[string]int{"foo": 1, "bar": 2} // inline initialization
	fmt.Println("map:", n)

	nums := []int{2, 3, 4}
	for i, num := range nums { // range iterates over various of data structures
		fmt.Println("index:", i, " num:" ,num )
	}

	keyValueMap := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range keyValueMap {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
