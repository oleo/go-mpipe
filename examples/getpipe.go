package main

import (
  "github.com/oleo/go-mpipe/mpipe"
	"fmt"
	"flag"
	"encoding/json"
)


func main() {

	boolPtr := flag.Bool("dump",false,"Dump config in JSON")
	flag.Parse()


	// Init towards redis
	mpipe.Init("redis:6379")
		if(*boolPtr) {
			fmt.Print("{\n  \"mpipes\": [\n")
		}
	for i,pipename := range mpipe.AvailablePipes() {
		mp :=  mpipe.Retrieve(pipename)
		if(*boolPtr) {
			// Dump registered available pipes as JSON
			b , err := json.MarshalIndent(mp," . ", "  ")
			if err != nil {
				fmt.Println(err)
				return
			}
			if(i==0) {
				fmt.Printf("%s",string(b))
			} else {
				fmt.Printf(",%s",string(b))
			}
		} else {
			// List registered available pipes
			fmt.Print("\n===================================================================\n")
			fmt.Printf("GOT %s \n",pipename)
			mpipe.Show(&mp)
		}
	}
	if(*boolPtr) {
			fmt.Print("] }")
	}

}

