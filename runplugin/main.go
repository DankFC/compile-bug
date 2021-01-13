package main

import "fmt"
import "plugin"

func main() {

	p, err := plugin.Open("/tmp/plugin.so")
	if err != nil {
		panic(err)
	}

	epoint, err := p.Lookup("Export")
	if err != nil {
		panic(err)
	}
	entrypoint, ok := epoint.(func())
	if !ok {
		fmt.Printf("Entrypoint doesn't have valid entry Definition\n")
		return
	}

	fmt.Printf("running plugin\n")
	entrypoint()
}
