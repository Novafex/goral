package fs

import (
	"encoding/json"
	"os"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

// Marshaler is a function that takes a target and converts it to bytes for
// data serialization
type Marshaler func(target any) ([]byte, error)

var (
	// marshalers is the list of known Marshaler functions based on the extension
	// type.
	marshalers = map[string]Marshaler{
		"yaml": yaml.Marshal,
		"toml": toml.Marshal,
		"json": json.Marshal,
	}
)

// Write takes an object and writes it to a path forming the filepath using the
// provided base and extension portions. The extension is separate so it knows
// which marshaler to use. If the extension is invalid, it will default to the
// first choice via [GetExtensionOrder].
func Write(obj any, base, ext string) error {
	path := CombineBaseExt(base, ext)

	marshaler, ok := marshalers[ext]
	if !ok {
		marshaler = marshalers[extensionOrder[0]]
	}

	data, err := marshaler(obj)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
