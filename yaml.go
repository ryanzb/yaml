package yaml

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Unmarshal(in []byte, out interface{}) (err error) {
	return yaml.Unmarshal(in, out)
}

func UnmarshalFile(path string, out interface{}) (err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, out)
}

func Marshal(in interface{}) (out []byte, err error) {
	return yaml.Marshal(in)
}