package main

import (
  "github.com/oleo/go-mpipe/mpipe"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)


func main() {

	// Init towards redis
	mpipe.Init("redis:6379")

	// Read file
	data, err :=  ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Print(err)
	}

	// parse mpipes
		var cfg mpipe.MPipeConfig


		err = json.Unmarshal(data, &cfg)
		if err != nil {
			fmt.Println("Error : " , err)
		}

		fmt.Println("Read configfile:")
		fmt.Printf("Site:      %s\n",cfg.Site)	
		fmt.Printf("ID:        %s\n",cfg.ID)	
		fmt.Printf("Nr. of pipes:  %d\n",len(cfg.MPipes))
	
		for _,currpipe := range cfg.MPipes {
			fmt.Printf(" -> Importing  : %10s  ... ", currpipe.Name)
			mpipe.Store(&currpipe)
			fmt.Println("Done")
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

