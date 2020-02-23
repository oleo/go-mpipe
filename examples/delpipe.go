package main

import (
  "github.com/oleo/go-mpipe/mpipe"
	"fmt"
	"os"
)


func main() {

	// Init towards redis
	mpipe.Init("redis:6379")

	// Read file
	pipename := os.Args[1]
		fmt.Printf("Deleting mpipe %s ...",pipename)
		mp := mpipe.Retrieve(pipename)
		if(len(mp.Name)>0) {
			mpipe.Delete(&mp)
			fmt.Println("Done")
		} else {
			fmt.Println("Error")
		}
		//Build MP
		/*var mp mpipe.MPipe
		mp.Name="My second pipe"
		mp.Connector="Redis"
		mp.Plugin="plainPipe"
		mp.ChannelIn=[]string{"first:34/topic/jalla","last:342/topic/fjosr"}
		mp.ChannelOut=[]string{"bad:334/topic/mainjalla"}

		//Store MP
		mpipe.storePipe(&mp2)

//
	fmt.Println("Updated config:")
	fmt.Println(mpipe.DumpJSONConfig())
	*/
}

