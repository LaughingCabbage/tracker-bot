package main

import (
	"io/ioutil"
)

//Key stores an API token
type Key struct {
	Value string
}

//LoadKey is used to load an API token from a file.
func LoadKey(path string) Key {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return Key{Value: string(data)}
}
