package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

var registeredFlags = make(map[string]any)

type Flags struct {
	Name  string
	Value string
}

func main() {
	// name := registerFlag("name")
	// age := registerFlag("age")

	// mymap := parseFlags()
	// getValues(mymap)
	// fmt.Print(name.Value)
	// fmt.Print(age.Value)
	fmt.Print(runtime.GOMAXPROCS(0))
}

func parseFlags() map[string]any {
	flags := os.Args[1:]
	flagMap := make(map[string]any)
	for _, flag := range flags {
		pair := strings.Split(flag, "=")
		flagMap[pair[0]] = pair[1]
	}
	return flagMap
}

func getValues(req map[string]any) {
	for k, v := range req {
		if pointer, ok := registeredFlags[k]; ok {
			if f, ok := pointer.(*Flags); ok {
				f.Value = v.(string)
			}
		}
	}
}

func registerFlag(name string) *Flags {
	f := &Flags{Name: name}
	registeredFlags[name] = f
	return f
}
