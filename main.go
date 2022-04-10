package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	ticker := time.NewTicker(5 * time.Minute)
	quit := make(chan struct{})

	// Print a Joke when runs
	getJoke()

	for {
		select {
		case <-ticker.C:
			// Get a random joke & notify
			getJoke()
		case <-quit:
			ticker.Stop()
			return
		}
	}

}

func getJoke() {
	// Chuck Norris Random Jokes api
	url := "https://api.chucknorris.io/jokes/random"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// Parse The Response
	var (
		data map[string]interface{}
	)
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}

	log.Println(data["value"])

	err = beeep.Notify("Chuck Norris Jokes!", data["value"].(string), "./assets/chuck.png")
	if err != nil {
		log.Fatal(err)
	}
}
