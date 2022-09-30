package cmd

import (
	"encoding/json"
	"strings"
)

func Slice2Map(slice []string) map[string]string {
	map1 := make(map[string]string)
	for _, v := range slice {
		s := strings.Split(v, "=")
		map1[s[0]] = s[1]
	}
	return map1
}

func Slice2JSON(slice []string) ([]byte, error) {
	map1 := Slice2Map(slice)
	jsonStr, err := json.Marshal(map1)
	if err != nil {
		return nil, err
	}
	return jsonStr, nil
}

func Map2JSON(map1 map[string][]string) ([]byte, error) {
	jsonStr, err := json.Marshal(map1)
	if err != nil {
		return nil, err
	}
	return jsonStr, nil
}

func Str2JSON(str string) ([]byte, error) {

	var map2 map[string]interface{}
	err := json.Unmarshal([]byte(json_input), &map2)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(&map2)
	if err != nil {
		return nil, err
	}

	return b, nil
}
