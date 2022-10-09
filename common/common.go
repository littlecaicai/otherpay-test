package common

import (
	"encoding/json"
)






func ConvToJSON(v interface{}) (str string) {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}
