package pkg

import (
	"strings"
	"testing"
)

func TestConvertToMapLiteral(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]string
		wantErr  bool
	}{
		{
			name:  "Valid JSON with string values",
			input: `{"name": "John", "city": "NY"}`,
			expected: map[string]string{
				"name": "'John'",
				"city": "'NY'",
			},
			wantErr: false,
		},
		{
			name:  "Valid JSON with non-string values",
			input: `{"age": 30, "height": 5.9, "isStudent": false}`,
			expected: map[string]string{
				"age":       "30",
				"height":    "5.9",
				"isStudent": "false",
			},
			wantErr: false,
		},
		{
			name:  "Valid JSON with mixed values",
			input: `{"name": "John", "age": 30, "isStudent": false}`,
			expected: map[string]string{
				"name":      "'John'",
				"age":       "30",
				"isStudent": "false",
			},
			wantErr: false,
		},
		{
			name:     "Empty JSON input",
			input:    `{}`,
			expected: map[string]string{},
			wantErr:  false,
		},
		{
			name:    "Invalid JSON input",
			input:   `{"name": "John"`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			got, err := ConvertToMapLiteral(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToMapLiteral() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !compareMaps(got, tt.expected) {
				t.Errorf("ConvertToMapLiteral() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// compareMaps is a helper function to compare two map[string]string
func compareMaps(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
