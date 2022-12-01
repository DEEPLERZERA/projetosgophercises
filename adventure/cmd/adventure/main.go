package main

import (
	"flag"
	"fmt"
	"os"

	"adventure"
)

func main() {
	filename := flag.String("file", "gopher.json", "Arquivo JSON com a história do Gopher")
	flag.Parse()
	fmt.Printf("Usando a história em %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := adventure.JsonStory(f)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)

}
