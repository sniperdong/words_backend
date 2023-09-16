package util

import "encoding/json"

func JsonEncode(data interface{}) string {
	s, _ := json.Marshal(data)
	return string(s)
}
