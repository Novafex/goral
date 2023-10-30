package fs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

// Unmarshaler is any function that can be used to unmarshal byte data from
// a file.
type Unmarshaler func(src []byte, target any) error

var (
	// unmarshalers is the list of known unmarshalers held by extension. These
	// extensions are the same from [GetExtensionOrder].
	unmarshalers = map[string]Unmarshaler{
		"yaml": yaml.Unmarshal,
		"toml": toml.Unmarshal,
		"json": json.Unmarshal,
	}
)

// Read takes a full path and the target interface and unmarshals the contents
// in that file using the unmarshaler detected by it's file extension.
func Read(path string, target any) error {
	// Clean and extract the file extension
	ext := CleanExtension(filepath.Ext(path))
	if !IsValidExtension(ext) {
		return fmt.Errorf("invalid extension %s", ext)
	}

	// Ensure we have an unmarshaler for this
	unmarshal, ok := unmarshalers[ext]
	if !ok {
		return fmt.Errorf("unexpected extension type %s", ext)
	}

	// Read the data into memory
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Do the magic
	return unmarshal(data, target)
}

// FindAndRead takes a base path to a file (without extension) and goes through
// the list of extensions in order of desire until one matches. When matched,
// it will Unmarshal the file's contents into the given target (interface) based
// on the extension that was found.
//
// If no file is found that matches any of the extensions a [os.ErrNotExist]
// error is returned.
func FindAndRead(base string, target any) error {
	// Find the file we want
	found, ext := FindPathWithExtensions(base)
	if !found {
		return os.ErrNotExist
	}

	return Read(CombineBaseExt(base, ext), target)
}
