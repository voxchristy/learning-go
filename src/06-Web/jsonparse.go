package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"math/big"
)

type Tour struct{
	Name, Price string
}

func main() {

	// sample test API
	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentfromserver(url)
	tours := toursfromjson(content)
	// fmt.Println(tours)

	for _, tour := range tours{
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
		fmt.Printf("%v ($%.2f)\n", tour.Name, price)
	} 
}

func checkerror(err error){
	if err != nil{
		panic(err)
	}
}


func contentfromserver(url string)string{
	res, err := http.Get(url)
	checkerror(err)

	// returns pointer to res obj
	fmt.Printf("Response type: %T\n: ", res)

	defer res.Body.Close()
	// read the content (in bytes)
	bytes, err := ioutil.ReadAll(res.Body)
	checkerror(err)

	// convert bytes to string and return data
	return string(bytes)

}


func toursfromjson(content string) []Tour{

	tours :=make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_,err:=decoder.Token()
	checkerror(err)

	var tour Tour
	for decoder.More(){
		err := decoder.Decode(&tour)
		checkerror(err)
		tours = append(tours, tour)
	}
	return tours

}
