package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

func StreamJsonFile(filename string, key string) (interface{}, error) {
	return streamJsonFile(filename, key)
}

func streamJsonFile(filename string, targetKey string) (interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	dec := json.NewDecoder(file)
	t, err := dec.Token()
	if err != nil || t != json.Delim('{') {
		return nil, fmt.Errorf("invalid json object: %w", err)
	}

	for dec.More() {
		tk, err := dec.Token()
		if err != nil {
			return nil, fmt.Errorf("failed to read token: %w", err)
		}
		key, ok := tk.(string)
		if !ok {
			continue
		}

		if key == targetKey {
			var value interface{}
			if err := dec.Decode(&value); err != nil {
				return nil, fmt.Errorf("failed to decode value: %w", err)
			}
			return value, nil
		}

		if err := skipValue(dec); err != nil {
			return nil, fmt.Errorf("failed to skip value: %w", err)
		}
	}
	return nil, nil

}

func skipValue(dec *json.Decoder) error {
	t, err := dec.Token()
	if err != nil {
		return err
	}

	switch delim := t.(type) {
	case json.Delim:
		var depth = 1
		open := delim
		cls := matchingDelim(open)
		for depth > 0 {
			t, err := dec.Token()
			if err != nil {
				return err
			}
			if d, ok := t.(json.Delim); ok {
				if d == open {
					depth++
				} else if d == cls {
					depth--
				}
			}
		}
	}
	return nil
}

func matchingDelim(d json.Delim) json.Delim {
	switch d {
	case '{':
		return '}'
	case '[':
		return ']'
	default:
		return 0
	}
}
