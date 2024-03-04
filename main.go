package main

import (
	"time"
	"log"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"math/rand"
	"github.com/tarikjaber/main/hello"
)

type Joke struct {
	Value string `json:"value"`
}

func getJoke(c chan string) {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	
	if err != nil {
		fmt.Println("Error with http request")	
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	
	if err != nil {
		fmt.Println("Error reading body")	
		os.Exit(1)
	}
	
	var joke Joke
	
	err = json.Unmarshal(body, &joke)
	if err != nil {
		fmt.Println("Error decoding json")
	}
	
	c <- joke.Value
}

func printMamaJoke() {
	c := make(chan string)

	go getJoke(c)
	fmt.Println(<-c)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file " + filePath, err)
	}
	defer f.Close()
	
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", filePath, err)
	}
	
	return records
}

func main() {
	rand.Seed(time.Now().UnixNano())
	Yoda()
	if rand.Intn(2) == 0 {
		records := readCsvFile("jokes.csv")
		randJokeIndex := rand.Intn(len(records))
		fmt.Println(records[randJokeIndex][1])
	} else {
		printMamaJoke()
	}
}