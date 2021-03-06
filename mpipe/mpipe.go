package mpipe

import (
  "strings"
	"fmt"
  "github.com/gomodule/redigo/redis"
	"encoding/json"
)

var (
	c redis.Conn
	err error
	reply interface{}
	redis_store = "redis:6379"
)

type MPipe struct {
	Name string
	Description string
  Connector string
  Plugin string
	ChannelIn []string
	ChannelOut []string
}

type MPipeConfig struct {
	Site string
  ID string
  MPipes []MPipe
}

func Init(redissrv string) {
	redis_store=redissrv
}

func Delete(mp *MPipe) {

	c, err = redis.Dial("tcp", redis_store)
	if err != nil {
		fmt.Print(err)
	}

	_, err = c.Do("JSON.DEL","mpipe."+mp.Name)
	if err != nil {
		fmt.Print(err)
	}

}

func Store(mp *MPipe) {
	b , err := json.Marshal(mp)
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err = redis.Dial("tcp", redis_store)
	if err != nil {
		fmt.Print(err)
	}

	_, err = c.Do("JSON.SET","mpipe."+mp.Name,".",string(b))
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

//	fmt.Printf(" Read %s struct\n",structdata.Name)
	return structdata

}

func AvailablePipes() []string {

	c, err = redis.Dial("tcp",redis_store) 
	if err != nil {
		fmt.Print(err)
	}

	keys, err := redis.Strings(c.Do("KEYS","mpipe.*"))
	if err != nil {
		fmt.Print(err)
	}

	var out []string

	for _, key :=  range keys {
		modified_key := strings.Replace(key,"mpipe.","",-1)
		out = append(out,modified_key)
	}
	/*
	for _, key :=  range keys {
		fmt.Printf("got %s\n",key)
	}
	*/
  return out
}

func DumpJSONConfig()  string {
	cfg := MPipeConfig{}
	cfg.Site="MySite"
	cfg.ID="Prod-1"

	for _,pipename := range AvailablePipes() {
			mp :=  Retrieve(pipename)
			cfg.MPipes = append(cfg.MPipes, mp)
	}
				// List registered available pipes
					b , err := json.MarshalIndent(cfg,"", "  ")
					if err != nil {
						return fmt.Sprintf("%v",err);
					} else {
						//fmt.Printf("%s",string(b))
						return string(b)
					}

}
