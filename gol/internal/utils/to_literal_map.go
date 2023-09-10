package utils

import (
	"fmt"
	json "github.com/json-iterator/go"
	"io"
)

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
