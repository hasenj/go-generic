package generic

import (
	"encoding/json"
	"strings"
	"os"
)

// ReadFromJSONFile fills in an object with data from the json provied by the
// file specified
func ReadFromJSONFile[T any](filepath string, obj *T) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(obj)
}

func JSONify(obj any, indent string) string {
	var buf strings.Builder
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", indent)
	enc.Encode(obj)
	return buf.String()
}