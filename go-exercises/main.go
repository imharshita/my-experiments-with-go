package main

import (
	"fmt"
	"log"
	"os"

	"github.com/harshita-sharma012/my-experiments-with-go/go-exercises/parse"
)

func main() {
	fd, err := os.Open("ex1.html")
	if err != nil {
		log.Fatal(err)
	}
	link, err := parse.Parse(fd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(link)
}
