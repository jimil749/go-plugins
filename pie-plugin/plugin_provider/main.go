package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/rpc/jsonrpc"

	utils "github.com/jimil749/pie-plugin/utils"
	"github.com/natefinch/pie"
)

func main() {
	log.SetPrefix("[plugin log] ")

	p := pie.NewProvider()
	if err := p.RegisterName("Plugin", api{}); err != nil {
		log.Fatalf("failed to register plugin: %s", err)
	}
	p.ServeCodec(jsonrpc.NewServerCodec)
}

type api struct{}

func (api) Put(args *utils.PutRequest, reply *utils.PutResponse) error {
	log.Printf("call for Put")
	value := []byte(fmt.Sprintf("%s", string(args.Value)))
	return ioutil.WriteFile("kv_"+args.Key, value, 0644)
}

func (api) Get(args *utils.GetRequest, reply *utils.GetResponse) error {
	log.Printf("call for get")
	var err error
	reply.Value, err = ioutil.ReadFile("kv_" + args.Key)
	if err != nil {
		log.Fatalf("Could read the key")
		return err
	}
	return nil
}
