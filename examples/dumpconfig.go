package main

import (
  "github.com/oleo/go-mpipe/mpipe"
	"fmt"
)


func main() {

	// Init towards redis
	mpipe.Init("redis:6379")

	fmt.Println(mpipe.DumpJSONConfig())

}

