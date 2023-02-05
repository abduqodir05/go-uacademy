package homeworks

import "errors"

func Drop(chars []byte, index int) (error, string, []byte) {
	message := ""
	data := []byte{}

	if len(chars) < index {
		return errors.New("smt"), "index more than length of array", data
	}
	for i := 0; i < len(chars); i++ {
		if i != index {
			data = append(data, chars[i])
		}

	}

	return nil, message, data
}
