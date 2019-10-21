package main

import (
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(generateRandomText(1000))

	}))
	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		log.Fatalln(err)
	}
}

func generateRandomText(n int) (text []byte) {
	for i := 0; i < n; i++ {
		text = strconv.AppendInt(text, rand.Int63()%10, 10)
	}

	return
}
