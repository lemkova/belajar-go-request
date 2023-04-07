package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	//Generate random number between 1 and 100 for variable wind and water with custom seed
	// Send the data to server with interval of 15 seconds with POST method to https://jsonplaceholder.typicode.com/posts
	for {
		rand.Seed(time.Now().UnixNano())
		wind := rand.Intn(100)
		water := rand.Intn(100)

		data := map[string]interface{}{
			"water": water,
			"wind":  wind,
		}

		requestJson, err := json.Marshal(data)
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(requestJson))
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(string(body))
		resp.Body.Close()

		var statusWind string
		if wind > 15 {
			statusWind = "Danger"
		} else if wind <= 15 && wind > 6 {
			statusWind = "Warning"
		} else {
			statusWind = "Safe"
		}
		fmt.Printf("Status Wind: %s\n", statusWind)

		var statusWater string
		if water > 8 {
			statusWater = "Danger"
		} else if water <= 8 && water > 6 {
			statusWater = "Warning"
		} else {
			statusWater = "Safe"
		}
		fmt.Printf("Status Water: %s\n", statusWater)

		time.Sleep(15 * time.Second)
	}

}
