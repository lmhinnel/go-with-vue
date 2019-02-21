package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/kansuke231/go-with-vue/api/models"
)

// function schedule spawns a go routine that updates StaticResult every delay specified.
func schedule(s *models.StaticResult, update func(s *models.StaticResult), delay time.Duration) chan bool {
	stop := make(chan bool)
	go func() {
		for {
			// Every delay (e.g. 5 minutes), update() gets executed.
			select {
			case <-time.After(delay):
				// Do update here.
				update(s)

			}
		}
	}()

	return stop
}

func updateStaticResult(s *models.StaticResult) {
	s.ID = rand.Int()
	s.SomeResult = "ThisIsSomeResult" + strconv.Itoa(rand.Int())
}

func generateStaticResults() *models.StaticResult {
	return &models.StaticResult{ID: rand.Int(), SomeResult: "ThisIsSomeResult" + strconv.Itoa(rand.Int())}
}
