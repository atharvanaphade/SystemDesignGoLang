package main

import "time"

var database map[string]string = map[string]string{
	"index.html": "<h1>Hello, World!</h1>",
}

func getData() string {
	time.Sleep(2 * time.Second)
	return database["index.html"]
}
