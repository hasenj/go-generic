package generic

import "os"
import "encoding/json"

func ReadFromFile[T any](filepath string, obj *T) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(obj)
}
