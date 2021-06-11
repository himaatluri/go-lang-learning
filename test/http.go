package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	Status := resp.Status
	request := resp.Request

	resp.Body.Close()
	fmt.Println(string(data), "\n")

	fmt.Println("\n", Status)
}
