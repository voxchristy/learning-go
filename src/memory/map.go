package main

// Map is an unordered collection of key-value pairs
// its essentially a Hashtable
import (
	"fmt"
	"sort"
)

func main() {

	//map declared
	states := make(map[string]string)
	fmt.Println(states)

	states["KL"] = "Kerala"
	states["MAH"] = "Maharsahtra"
	states["GA"] = "Goa"
	states["DL"] = "New Delhi"
	fmt.Println(states)

	//read from map
	goc := states["KL"]
	fmt.Println("Gods own country: ", goc)

	//delete
	delete(states, "GA")
	fmt.Println(states)

	//looping in a map
	for k, v := range states {
		fmt.Printf("%v - %v \n", k, v)
	}

	//sorting
	// make(type, init-size)
	keys := make([]string, len(states))
	fmt.Println("Keys: ", keys)

	i := 0
	for k := range states {
		// add keys to map
		keys[i] = k
		i++
	}
	fmt.Println("Keys: ", keys)

	// sort the keys
	sort.Strings(keys)
	fmt.Println("*****SORTED******")
	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}
