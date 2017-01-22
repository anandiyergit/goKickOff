package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	states := make(map[string]string)
// 	fmt.Println(states)

// 	states["MH"] = "Maharashtra"
// 	states["AP"] = "Andhra Pradesh"
// 	states["MP"] = "Madhya Pradesh"
// 	fmt.Println(states)

// 	maharashtra := states["MH"]
// 	fmt.Println(maharashtra)

// 	delete(states, "MP")
// 	fmt.Println(states)

// 	states["GJ"] = "Gujrat"

// 	for k, v := range states {
// 		fmt.Printf("%v: %v\n", k, v)
// 	}

// 	keys := make([]string, len(states))
// 	i := 0
// 	for k := range states {
// 		keys[i] = k
// 		i++
// 	}
// 	sort.Strings(keys)
// 	fmt.Println("\nSorted")
// 	for i := range keys {
// 		fmt.Println(states[keys[i]])
// 	}

// }
