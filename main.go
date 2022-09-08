package main

import (
	"net/http"
	"os"

	"quiiLearning/countdown"
	"quiiLearning/greetings"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	greetings.Greet(w, "world")
}

func ServeTheGreaterHandler() {
	http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler))
}

func main() {
	countdown.Countdown(os.Stdout, &countdown.RealSleeper{})
}
