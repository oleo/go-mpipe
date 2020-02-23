package mpipe

import (
	"fmt"
  "github.com/gomodule/redigo/redis"
	"encoding/json"
	"os"
)

var (
	c redis.Conn
	err error
	reply interface{}
)
var redis_store string "redis:6379"

type MPipe struct {
	Name string
  Connector string
  Plugin string
	ChannelIn []string
	ChannelOut []string
}

func Init(redissrv string) {
	redis_store=redissrv
}
func Store(mp *MPipe) {
	b , err := json.Marshal(mp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Will store:")
	fmt.Println(string(b))

  // Handle pooling using several redis clients... //

	c, err = redis.Dial("tcp", redis_store)
	if err != nil {
		fmt.Print(err)
	}

	_, err = c.Do("JSON.SET","mpipe.first",".",string(b))
	if err != nil {
		fmt.Print(err)
	}

}

func Show(mp *MPipe) {
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
func Retrieve(pipename string) MPipe {
  // Handle pooling using several redis clients... //

	c, err = redis.Dial("tcp",redis_store) 
	if err != nil {
		fmt.Print(err)
	}

	jsondata, err := redis.String(c.Do("JSON.GET","mpipe."+pipename))
	if err != nil {
		fmt.Print(err)
	}

	structdata :=  MPipe{}
  _ = json.Unmarshal([]byte(jsondata),&structdata)

	fmt.Printf(" Read %s struct\n",structdata.Name)
	return structdata

}
