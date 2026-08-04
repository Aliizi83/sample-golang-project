package common

import "encoding/json"

func TypeConverter[T any](data any) (*T, error) {
	var result = new(T)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(jsonData, &result)
	return result, err
}
