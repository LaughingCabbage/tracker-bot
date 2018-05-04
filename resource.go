package main

//Key stores an API token
type Key struct {
	Value string
}

//LoadKey is used to load an API token from a file.
func LoadKey() Key {
	return Key{}
}
