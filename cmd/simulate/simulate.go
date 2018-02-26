package main

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"strconv"
	"sync"
	"fmt"
	"time"
	"log"
)

func trackTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	defer trackTime(time.Now(), "main()")
	url := "https://azj3z8mlq6.execute-api.eu-west-1.amazonaws.com/Prod/"
	cardOrder := []byte(`{"order": ["maki-1", "maki-2", "maki-3", "sashimi", "egg", "salmon", "squid", "wasabi", "pudding", "tempura", "dumpling", "tofu", "eel", "temaki"]}`)
	nRequests := 1000

	var wg sync.WaitGroup
	wg.Add(nRequests)
	results := make([]int, nRequests)

	for i := 0; i < nRequests; i++ {
		go func(url string, cardOrder []byte, i int) {
			defer wg.Done()
			request := bytes.NewBuffer(cardOrder)
			result, err := http.Post(url, "application/json", request)
			if err != nil {
				fmt.Println("Error: %v", err)
			} else {
				body, err := ioutil.ReadAll(result.Body)
				if err != nil {
					fmt.Println(err)
				} else {
					r, _ := strconv.Atoi(string(body))
					results[i] = r
				}
			}
		}(url, cardOrder, i)
	}

	wg.Wait()

	log.Printf("Executed %v simulations, result = %v", nRequests, sum(results...))
}
