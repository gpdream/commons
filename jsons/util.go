package jsons

import "encoding/json"

func ToJson(v interface{}) (string,error) {
	data,err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func ParseJson(jsonStr string, v interface{}) error {
	bytes := []byte(jsonStr)
	return json.Unmarshal(bytes, v)
}
