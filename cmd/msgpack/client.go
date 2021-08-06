package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	msgpack "github.com/vmihailenco/msgpack/v5"
)

func post() {
	fmt.Println("Post begin")

	url := "http://localhost:9090/todos"
	method := "PUT"

	msg, err := msgpack.Marshal(&Task{Task: "task1"})
	if err != nil {
		log.Panic(err)
	}

	payload := bytes.NewReader(msg)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Add("Content-Type", "application/msgpack")

	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(body))

	fmt.Println("Post end")
}

func get() {
	fmt.Println("Get begin")

	url := "http://localhost:9090/todos"
	method := "GET"

	// msg, err := msgpack.Marshal(&Task{Task: "task1"})
	// if err != nil {
	// 	log.Panic(err)
	// }

	// payload := bytes.NewReader(msg)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Panic(err)
	}
	// req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(body))

	fmt.Println("Get end")
}

func main() {
	post()
	get()
}

// Generated by https://quicktype.io

type Task struct {
	Task string `json:"task"`
}
