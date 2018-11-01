package qstruct

import (
	"encoding/json"
	"net/url"
)

func Decode(u url.Values, v interface{}) error {
	var val = make(map[string]interface{})
	for k, _ := range u {
		val[k] = u.Get(k)
	}

	j, err := json.Marshal(val)
	if err != nil {
		return err
	}

	return json.Unmarshal(j, v)
}
