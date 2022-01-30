package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	fileInput,err := os.Open("./Index.html")
	if err != nil {
		panic(err)
	}

	defer func ()  {
		if err:= fileInput.Close(); err !=nil {
			panic(err)
		}
	}

}
