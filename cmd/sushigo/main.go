package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"log"
	"os/signal"
	"os"
	"context"
	"html/template"
	"sync"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	port = kingpin.Flag("port", "Port number").Short('p').Default("8080").Int()
)

type SushiGoCard struct {
	Title   string
	ImgPath string
}

type PageData struct {
	SushiGoCards []SushiGoCard
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	pageData := PageData{
		SushiGoCards: []SushiGoCard{
			{"Chopsticks", "/static/img/chopsticks.png"},
			{"Dumpling", "/static/img/dumpling.png"},
		},
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, pageData)
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func simulateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	input := r.Form["lambdaInput"][0]

	url := "https://azj3z8mlq6.execute-api.eu-west-1.amazonaws.com/Prod/"
	nRequests := 100
	var wg sync.WaitGroup
	wg.Add(nRequests)
	results := make([]int, nRequests)
	failures := 0

	start := time.Now()
	for i := 0; i < nRequests; i++ {
		go func(url string, cardOrder string, i int) {
			defer wg.Done()
			result, err := http.Post(url, "application/json", strings.NewReader(cardOrder))
			if err != nil {
				fmt.Printf("Error: %v", err)
				failures += 1
			} else {
				body, err := ioutil.ReadAll(result.Body)
				if err != nil {
					fmt.Println(err)
					failures += 1
				} else {
					r, _ := strconv.Atoi(string(body))
					results[i] = r
				}
			}
		}(url, input, i)
	}

	wg.Wait()
	elapsed := time.Since(start)
	w.Write([]byte(fmt.Sprintf(
		"Executed %v simulations. "+
			"Result = %v. "+
			"Failures = %v. "+
			"Duration = %v.",
		nRequests, sum(results...), failures, elapsed,
	)))
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	kingpin.Parse()

	r := mux.NewRouter()
	r.Use(logRequests)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/simulate", simulateHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%v", *port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Graceful shutdown. See https://github.com/gorilla/mux#graceful-shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Sushi Go shutting down.")
	os.Exit(0)
}
