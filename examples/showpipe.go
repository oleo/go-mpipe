package main

import (
	"fmt"
  "github.com/gomodule/redigo/redis"
	"encoding/json"
)

var (
	c redis.Conn
	err error
	reply interface{}
)

type MPipe struct {
	Name string
  Connector string
  Plugin string
	ChannelIn []string
	ChannelOut []string
}

func main() {


	var mp MPipe
	mp.Name="My brand new piped"
	mp.Connector="Redis"
	mp.Plugin="plainPipe"
	mp.ChannelIn=[]string{"first:34/topic/jalla","last:342/topic/fjosr"}
	mp.ChannelOut=[]string{"bad:334/topic/mainjalla"}
	
	showPipe(&mp)
	storePipe(&mp)
	var mp2 MPipe
	mp2.Name="My second pipe"
	mp2.Connector="Redis"
	mp2.Plugin="plainPipe"
	mp2.ChannelIn=[]string{"first:34/topic/jalla","last:342/topic/fjosr"}
	mp2.ChannelOut=[]string{"bad:334/topic/mainjalla"}
	
	showPipe(&mp2)
	storePipe(&mp2)

	//myp := &MPipe{Name: "test",Connector:"jalla"}

}

func storePipe(mp *MPipe) {
	b , err := json.Marshal(mp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Will store:")
	fmt.Println(string(b))

  // Handle pooling using several redis clients... //

	c, err = redis.Dial("tcp", "redis:6379")
	if err != nil {
		fmt.Print(err)
	}

	_, err = c.Do("JSON.SET","mpipe.second",".",string(b))
	if err != nil {
		fmt.Print(err)
	}

}

func showPipe(mp *MPipe) {
	fmt.Printf("\nMPipeInfo:\n")
	fmt.Print("--------------------------------------------------------\n")
	fmt.Printf("Name:       			%s\n",mp.Name);
	fmt.Printf("Connector:  			%s\n",mp.Connector);
	fmt.Printf("Plugin:     			%s\n",mp.Plugin);

	fmt.Print("\nThe following channels are defined:\n\n")
	//fmt.Printf("ChannelIn has %d items:\n",len(mp.channelIn))
	for _, element := range mp.ChannelIn {
		fmt.Printf("  %30s ---> \n",element)
	}
	//fmt.Printf("ChannelOut has %d items:\n",len(mp.channelOut))
	for _, element := range mp.ChannelOut {
		fmt.Printf("  %30s ---> %s \n","",element)
	}
	fmt.Print("--------------------------------------------------------\n")
}
