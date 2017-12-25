package main

import (
	"io"
	"net/http"
	"log"
	"math/rand"
	"time"
	"fmt"
	"strconv"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	fmt.Println(r.Intn(10))
	response := "Hello http server by golang!\n"
	response += "now: " + time.Now().String() + "\n"
	response += "rand: " + strconv.Itoa(r.Intn(10)) + "\n"

	io.WriteString(w, response)
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12345", nil)

	if err != nil {
		log.Fatal("ListenAndServe fail: ", err)
	}
}
