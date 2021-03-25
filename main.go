package main

// Here we will import the kv plugin that we built.
import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
)

func main() {
	plugins, err := filepath.Glob("kv.so")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nLoading plugin %s\n", plugins[0])
	p, err := plugin.Open(plugins[0])
	if err != nil {
		panic(err)
	}

	symbol, err := p.Lookup("Put")
	if err != nil {
		panic(err)
	}

	putFunc, ok := symbol.(func(string, []byte) error)
	if !ok {
		panic("\nPlugin has no put")
	}

	err = putFunc("master", []byte("server"))
	if err != nil {
		fmt.Printf("Error in put")
		os.Exit(1)
	}

	symbol, err = p.Lookup("Get")
	getFunc, ok := symbol.(func(string) ([]byte, error))
	if !ok {
		panic("\nPlugin has no get")
	}
	val, err := getFunc("master")
	if err != nil {
		fmt.Printf("Error in get")
		os.Exit(1)
	}
	fmt.Printf("\nValue is : %s", string(val))
}
