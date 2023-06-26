package utils

import "github.com/whioue/cloud-go-sdk-V1.0/pkg/json"

type MapStringInterface map[string]string

func ConvertStruct2Map(v interface{}) ([]byte, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
