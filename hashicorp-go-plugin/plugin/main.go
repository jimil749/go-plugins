package main

import (
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/go-plugin"
	"github.com/jimil749/hashicorp-go-plugin/shared"
)

// KV is the implementation of KV. Real KV operations are performed here
type KV struct{}

func (KV) Put(key string, value []byte) error {
	value = []byte(fmt.Sprintf("%s\n\nWritten from plugin", string(value)))
	return ioutil.WriteFile("kv_"+key, value, 0644)
}

func (KV) Get(key string) ([]byte, error) {
	return ioutil.ReadFile("kv_" + key)
}

// Serves the plugin for the host process.
func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"kv": &shared.KVGRPCPlugin{Impl: &KV{}},
		},

		GRPCServer: plugin.DefaultGRPCServer,
	})
}
