package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Pogo ...
type Pogo struct {
	Msg string `json:"msg"`
}

//WeatherURL ...
const WeatherURL = "http://rest.shaoke.me/api/v1/weather/"

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		var city string
		var url string
		log.Print("Enter the city below(Press ctrl+c to exit):")
		city, _ = reader.ReadString('\n')
		log.Print("city:", city)
		url = WeatherURL + city
		weatherService(url)
	}

}

func weatherService(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode == 200 {
		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		info := Pogo{}
		json.Unmarshal(bytes, &info)
		log.Println("Content:", info.Msg)
	} else {
		log.Println("Not a valid city ! \n The status code is:", res.StatusCode)
	}
}
