package nginxproxymanager

import (
	"encoding/json"
)

type boolAsInt bool

func (b *boolAsInt) UnmarshalJSON(data []byte) error {
	var i int
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	*b = boolAsInt(i != 0)
	return nil
}

func (b boolAsInt) MarshalJSON() ([]byte, error) {
	if b {
		return json.Marshal(1)
	}
	return json.Marshal(0)
}

func (b boolAsInt) Bool() bool {
	return bool(b)
}
