package cliflags

import (
	"os"
	"strings"
)

var registeredFlags = make(map[string]any)

type flags struct {
	Name  string
	Value string
}

// func main() {
// 	name := RegisterFlag("name")
// 	age := RegisterFlag("age")

// 	ParseFlags()
// 	fmt.Print(name.Value)
// 	fmt.Print(age.Value)

// }

func ParseFlags() {
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
			if f, ok := pointer.(*flags); ok {
				f.Value = v.(string)
			}
		}
	}
}

func RegisterFlag(name string) *flags {
	f := &flags{Name: name}
	registeredFlags[name] = f
	return f
}
