package main

import (
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"

	utils "github.com/jimil749/pie-plugin/utils"
	"github.com/natefinch/pie"
)

func main() {
	log.SetPrefix("[master log] ")

	path := "plugin_provider"

	client, err := pie.StartProviderCodec(jsonrpc.NewClientCodec, os.Stderr, path)
	if err != nil {
		log.Fatalf("Error running plugin: %s", err)
	}
	defer client.Close()
	p := plug{client}
	res, err := p.Put("master", []byte("Server"))
	if err != nil {
		log.Fatalf("error calling Put: %s", err)
	}
	log.Printf("Response from plugin: %q", res)

	value, err := p.Get("master")
	if err != nil {
		log.Fatalf("error calling Get: %s", err)
	}
	log.Printf("Response from plugin: %q", value.Value)
}

type plug struct {
	client *rpc.Client
}

func (p plug) Put(key string, value []byte) (result *utils.PutResponse, err error) {
	args := &utils.PutRequest{
		Key:   key,
		Value: value,
	}
	err = p.client.Call("Plugin.Put", args, &result)
	return result, err
}

func (p plug) Get(key string) (result *utils.GetResponse, err error) {
	args := &utils.PutRequest{
		Key: key,
	}
	err = p.client.Call("Plugin.Get", args, &result)
	return result, err
}
