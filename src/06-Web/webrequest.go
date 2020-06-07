package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "http://api.plos.org/search?q=title:DNA"

	res, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	// returns pointer to res obj
	fmt.Printf("Response type: %T\n: ", res)

	defer res.Body.Close()

	// read the content (in bytes)
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil{
		panic(err)
	}

	// convert bytes to string
	content := string(bytes)
	fmt.Print(content)

}
