package pkg

import (
	"fmt"
	json "github.com/json-iterator/go"
	"io"
)

// ConvertToMapLiteral converts a JSON object to a map[string]string, where the
// values are converted to string literals.
// e.g. {"name": "John", "age": 30} -> {"name": "'John'", "age": "30"}
func ConvertToMapLiteral(r io.Reader) (map[string]string, error) {
	var inputData map[string]interface{}
	if err := json.NewDecoder(r).Decode(&inputData); err != nil {
		return nil, err
	}
	outputData := make(map[string]string)
	for key, value := range inputData {
		switch v := value.(type) {
		case string:
			outputData[key] = `'` + v + `'`
		default:
			outputData[key] = fmt.Sprintf("%v", v)
		}
	}
	return outputData, nil
}
