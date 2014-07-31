package extract

import (
	"encoding/json"
	"errors"
)

const (
	comma     = 44
	obrace    = 123
	cbrace    = 125
	obracket  = 91
	cbracket  = 93
	colon     = 58
	mark      = 34
	backslash = 92
)

func Extract(buf []byte, key string) (interface{}, error) {
	inString := false
	isKey := true
	level := 0
	chars := []byte(key)

	for i := 0; i < len(buf); i++ {
		c := buf[i]

		if c == backslash {
			i++
			continue
		}

		if c == mark {
			inString = !inString
			continue
		}

		if !inString {
			if c == colon {
				isKey = false
			} else if c == comma {
				isKey = true
			} else if c == obrace {
				level++
			} else if c == cbrace {
				level--
			}
		}
		if !isKey || level > 1 {
			continue
		}

		if isMatch(buf, i, chars) {
			start := i + len(key) + 2
			end, err := findEnd(buf, start)
			if err != nil {
				return nil, err
			}
			return parse(buf[start:end])
		}
	}

	return nil, errors.New("key not found")
}

func isMatch(buf []byte, i int, chars []byte) bool {
	if i > 0 && buf[i-1] != mark || len(buf) < i+len(chars) {
		return false
	}
	for j := 0; j < len(chars); j++ {
		if buf[i+j] != chars[j] {
			return false
		}
	}
	if buf[i+len(chars)] != mark {
		return false
	}
	return true
}

func findEnd(buf []byte, start int) (int, error) {
	if len(buf) <= start {
		return -1, errors.New("json too short")
	}

	level := 0
	s := buf[start]

	for i := start; i < len(buf); i++ {
		c := buf[i]
		if c == obrace || c == obracket {
			level++
			continue
		} else if c == cbrace || c == cbracket {
			level--
			if level > 0 {
				continue
			}
		}
		if level < 0 || level == 0 && (c == comma || c == cbrace || c == cbracket) {
			if s == obrace || s == obracket {
				return i + 1, nil
			} else {
				return i, nil
			}
		}
	}

	return -1, errors.New("missing end")
}

func parse(slice []byte) (interface{}, error) {
	var ret interface{}
	err := json.Unmarshal(slice, &ret)
	return ret, err
}
