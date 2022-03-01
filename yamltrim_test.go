package yamltrim

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestScalarYamlTrim(t *testing.T) {
	tests := map[string]struct {
		input    interface{}
		expected interface{}
	}{
		"scalar nil": {
			input:    nil,
			expected: nil,
		},
		"scalar string": {
			input:    "string",
			expected: "string",
		},
		"empty string": {
			input:    "",
			expected: nil,
		},
		"nonzero int": {
			input:    42,
			expected: 42,
		},
		"zero int": {
			input:    0,
			expected: nil,
		},
		"nonzero float": {
			input:    3.14,
			expected: 3.14,
		},
		"zero float": {
			input:    0.0,
			expected: nil,
		},
		"true boolean": {
			input:    true,
			expected: true,
		},
		"false boolean": {
			input:    false,
			expected: nil,
		},
		"empty slice": {
			input:    []interface{}{},
			expected: nil,
		},
		"nonempty slice": {
			input:    []interface{}{"one", "two"},
			expected: []interface{}{"one", "two"},
		},
		"reduce slice": {
			input:    []interface{}{"", "one", 0, "two", 0.0},
			expected: []interface{}{"one", "two"},
		},
		"reduce empty slice": {
			input:    []interface{}{"", "", ""},
			expected: nil,
		},
		"empty map": {
			input:    map[string]interface{}{},
			expected: nil,
		},
		"nonempty map": {
			input:    map[string]interface{}{"a": "one", "b": 42},
			expected: map[string]interface{}{"a": "one", "b": 42},
		},
		"reduce map": {
			input:    map[string]interface{}{"removeme": "", "a": "one", "b": 42, "c": 0.0},
			expected: map[string]interface{}{"a": "one", "b": 42},
		},
		"reduce empty map": {
			input:    map[string]interface{}{"removeme": "", "a": 0, "b": "", "c": 0.0},
			expected: nil,
		},
	}
	for description, test := range tests {
		result := YamlTrim(test.input)
		assert.Equal(t, test.expected, result, description)
	}
}

func TestYamlYamlTrim(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"simple scalar": {
			input:    `a: b`,
			expected: `a: b`,
		},
		"empty scalar": {
			input:    `a: ""`,
			expected: ``,
		},
		"simple list": {
			input: `
a:
  - 42
`,
			expected: `a: [ 42 ]`,
		},
		"empty list": {
			input: `
a:
  - 0
  - ""
`,
			expected: ``,
		},
		"deep structure": {
			input: `
top:
  middle:
    deep: value
  other:
  - one
  - two
`,
			expected: `top: {middle: {deep: value}, other: [one, two]}`,
		},
		"partial structure reduction": {
			input: `
top:
  middle:
    deep: ""
  other:
  - ""
  - two
`,
			expected: `top: {other: [two]}`,
		},
		"full structure reduction": {
			input: `
top:
  middle:
    deep: ""
  other:
  - ""
  - nested:
    very:
    - deep:
      - 0.0
`,
			expected: ``,
		},
	}
	for description, test := range tests {
		var input, expected interface{}
		err := yaml.Unmarshal([]byte(test.input), &input)
		assert.NoError(t, err)
		err = yaml.Unmarshal([]byte(test.expected), &expected)
		assert.NoError(t, err)
		result := YamlTrim(input)
		assert.Equal(t, expected, result, description)
	}
}
