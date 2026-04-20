package cliFlags

import (
	"os"
	"strings"
)

var registeredFlags = make(map[string]any)

type Flags struct {
	Name  string
	Value string
}

// func main() {
// 	name := registerFlag("name")
// 	age := registerFlag("age")

// 	parseFlags()
// 	fmt.Print(name.Value)
// 	fmt.Print(age.Value)

// }

func parseFlags() {
	flags := os.Args[1:]
	flagMap := make(map[string]any)
	for _, flag := range flags {
		pair := strings.Split(flag, "=")
		flagMap[pair[0]] = pair[1]
	}
	getValues(flagMap)
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
